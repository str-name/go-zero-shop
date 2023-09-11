package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/order/db/model"
	"zero-shop/payment/rpc/payment"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"

	"zero-shop/order/rpc/internal/svc"
	"zero-shop/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	WalletPay = "平台钱包支付"
	WxPay     = "微信支付"
	AlliPay   = "支付宝支付"
)

type OrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDetailLogic) OrderDetail(in *pb.OrderDetailReq) (*pb.OrderDetailResp, error) {
	// todo: add your logic here and delete this line

	// 获取订单信息
	var order model.Order
	err := l.svcCtx.OrderDB.Where("user_id = ? and order_sn = ? and del_state = ?", in.UserID, in.OrderSn, globalKey.DelStateNo).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_EXISTS_ERROR), "UserID: %v, OrderSn: %v", in.UserID, in.OrderSn)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE order ERROR: %+v", err)
	}
	// 获取商品信息
	productResp, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &product.ProductDetailReq{ProductID: order.ProductID})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Order Use ProductRpc.ProductDetail Fail"),
			"Order Use ProductRpc.ProductDetail Fail, productID: %v, ERROR: %+v", order.ProductID, err)
	}
	firstBanner := strings.Split(productResp.Product.Banner, ",")[0]
	// 获取收货地址信息
	addrResp, err := l.svcCtx.UserRpc.GetUserAddressDetail(l.ctx, &user.GetUserAddressDetailReq{AddressID: order.UserAddressID})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Order Use UserRpc.GetUserAddressDetail Fail"),
			"Order Use UserRpc.GetUserAddressDetail Fail, addressID: %v, ERROR: %+v", order.UserAddressID, err)
	}
	addrDetail := fmt.Sprintf("%s,%s,%s,%s", addrResp.Province, addrResp.City, addrResp.Region, addrResp.DetailAddress)
	addrPN := fmt.Sprintf("%s,%s", addrResp.Name, addrResp.Phone)

	var res = pb.OrderDetailResp{
		ID:               order.ID,
		CreateTime:       order.CreateTime.Format("2006-04-05 15:04:05"),
		UpdateTime:       order.CreateTime.Format("2006-04-05 15:04:05"),
		OrderSn:          order.OrderSn,
		UserID:           order.UserID,
		AddressDetail:    addrDetail,
		AddressPhoneName: addrPN,
		ProductID:        order.ProductID,
		Title:            productResp.Product.Title,
		SubTitle:         productResp.Product.SubTitle,
		Banner:           firstBanner,
		Info:             productResp.Product.Introduction,
		ProductStoreID:   order.ProductStoreID,
		ProductBossID:    order.ProductBossID,
		ProductCount:     order.ProductCount,
		UnitPrice:        order.UnitPrice,
		TotalPrice:       order.TotalPrice,
		Status:           order.Status,
		Remark:           order.Remark,
		PayTime:          "",
		PayType:          "",
	}

	if order.Status >= globalKey.OrderPayed {
		// 已经支付过会有支付信息
		payResp, err := l.svcCtx.PaymentRpc.GetPaymentDetail(l.ctx, &payment.GetPaymentDetailReq{OrderSn: order.OrderSn})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("Order Use PaymentRpc.GetPaymentDetail Fail"),
				"Order Use PaymentRpc.GetPaymentDetail Fail, orderSn: %v, ERROR: %+v", order.OrderSn, err)
		}
		switch payResp.PayMode {
		case globalKey.WalletPay:
			res.PayType = WalletPay
		case globalKey.WxPay:
			res.PayType = WxPay
		case globalKey.AlliPay:
			res.PayType = AlliPay
		default:
			res.PayType = WalletPay
		}
		res.PayTime = payResp.PayTime
	}

	return &res, nil
}
