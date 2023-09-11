package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
	"zero-shop/common/globalKey"
	"zero-shop/common/kqOrder"
	"zero-shop/common/unique"
	"zero-shop/common/xerr"
	model2 "zero-shop/order/db/model"
	"zero-shop/order/rpc/order"
	"zero-shop/payment/db/model"
	"zero-shop/user/rpc/user"

	"zero-shop/payment/rpc/internal/svc"
	"zero-shop/payment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentLogic {
	return &OrderPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPaymentLogic) OrderPayment(in *pb.OrderPaymentReq) (*pb.OrderPaymentResp, error) {
	// todo: add your logic here and delete this line

	// 判断订单是否存在
	orderResp, err := l.svcCtx.OrderRpc.GetOrderOnlyDetail(l.ctx, &order.GetOrderOnlyDetailReq{
		UserID:  in.UserID,
		OrderSn: in.OrderSn,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("PaymentRpc USE OrderRpc ERROR"),
			"PaymentRpc USE OrderRpc ERROR, userID: %v, orderSn: %v, err: %+v", in.UserID, in.OrderSn, err)
	}

	// 判断订单状态是否为未支付
	if orderResp.Status != globalKey.OrderWaitPay {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_ORDER_STATUS_ERROR),
			"userID: %v, orderSn: %v, price: %v", in.UserID, in.OrderSn, orderResp.TotalPrice)
	}

	var order = model2.Order{
		ID:             orderResp.ID,
		OrderSn:        orderResp.OrderSn,
		UserID:         orderResp.UserID,
		ProductID:      orderResp.ProductID,
		ProductStoreID: orderResp.ProductStoreID,
		ProductBossID:  orderResp.ProductBossID,
		ProductCount:   orderResp.ProductCount,
		UnitPrice:      orderResp.UnitPrice,
		TotalPrice:     orderResp.TotalPrice,
		Status:         orderResp.Status,
		Remark:         orderResp.Remark,
	}

	var paySn string
	switch in.ServiceType {
	case globalKey.WxPay:
		paySn, err = l.wxPay(order, in.UserID)
	case globalKey.AlliPay:
		paySn, err = l.alliPay(order, in.UserID)
	case globalKey.WalletPay:
		paySn, err = l.walletPay(order, in.UserID)
		fmt.Println("2222")
	}
	if err != nil {
		return nil, err
	}

	return &pb.OrderPaymentResp{
		PayTotalPrice: order.TotalPrice,
		PaySn:         paySn,
	}, nil
}

func (l *OrderPaymentLogic) walletPay(order model2.Order, userID int64) (string, error) {

	// 获取用户信息
	userResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{UserID: userID})
	if err != nil {
		return "", errors.Wrapf(xerr.NewErrMsg("PaymentRpc`walletPay USE UserRpc`GetUserInfo ERROR"),
			"PaymentRpc`walletPay USE UserRpc`GetUserInfo ERROR: %+v, userID: %v", err, userID)
	}
	// 判断用户id是否正确
	if userID != order.UserID {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_USER_ERROR), "orderSn: %v, userID: %v", order.OrderSn, userID)
	}
	// 判断钱包的金额是否够支付订单的金额
	if userResp.Money < order.TotalPrice {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_MONEY_NOT_ENOUGH_ERROR),
			"user`Money: %v, order`totalPrice: %v", userResp.Money, order.TotalPrice)
	}

	// 生成支付流水
	paymentSn := unique.GenerateSn(unique.PAYMENT_PREFIX)
	err = l.svcCtx.PaymentDB.Transaction(func(tx *gorm.DB) error {

		// 生成kq消息
		kqMessage, err := json.Marshal(kqOrder.PaymentUpdateOrderState{
			OrderSn:    order.OrderSn,
			UserID:     userID,
			OrderState: globalKey.OrderPayed,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("PaymentRPC USE MqKqOrder Generate message ERROR"),
				"PaymentRPC USE MqKqOrder Generate message ERROR: %+v", err)
		}

		// 如果金额足够，更新用户金额信息
		_, err = l.svcCtx.UserRpc.UpdateUserMoney(l.ctx, &user.UpdateUserMoneyReq{
			UserID: userID,
			Money:  userResp.Money - order.TotalPrice,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("PaymentRpc USE UserRpc.UpdateUserMoney ERROR"),
				"PaymentRpc USE UserRpc.UpdateUserMoney ERROR: %+v", err)
		}
		// 创建支付信息
		err = tx.Create(&model.Payment{
			PaymentSn:      paymentSn,
			OrderSn:        order.OrderSn,
			UserID:         userID,
			PayMode:        globalKey.WalletPay,
			PayTotal:       order.TotalPrice,
			TradeStateDesc: order.Remark,
			PayStatus:      globalKey.PaySuccess,
			PayTime:        time.Now(),
		}).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE payment ERROR: %+v", err)
		}

		// 通知修改订单信息，由未付款改为已付款（使用kafka消息队列）
		err = l.svcCtx.MqKqOrder.Push(string(kqMessage))
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("PaymentRPC USE MqKqOrder.Push ERROR"), "PaymentRPC USE MqKqOrder.Push ERROR: %+v", err)
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return paymentSn, nil
}

func (l *OrderPaymentLogic) wxPay(order model2.Order, userID int64) (string, error) {

	return "", nil
}

func (l *OrderPaymentLogic) alliPay(order model2.Order, userID int64) (string, error) {

	return "", nil
}
