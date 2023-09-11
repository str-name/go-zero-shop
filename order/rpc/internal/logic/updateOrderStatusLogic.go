package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/order/db/model"

	"zero-shop/order/rpc/internal/svc"
	"zero-shop/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	// todo: add your logic here and delete this line

	var order model.Order
	err := l.svcCtx.OrderDB.Where("user_id = ? and order_sn = ? and del_state = ?", in.UserID, in.OrderSn, globalKey.DelStateNo).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_EXISTS_ERROR), "UserID: %v, OrderSn: %v", in.UserID, in.OrderSn)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE order ERROR: %+v", err)
	}

	// 检查订单状态修改是否合理
	if err := l.verifyOrderStatus(in.OrderStatus, order.Status); err != nil {
		return nil, err
	}

	err = l.svcCtx.OrderDB.Model(&order).Updates(map[string]interface{}{
		"status": in.OrderStatus,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATE order ERROR: %+v", err)
	}

	return &pb.UpdateOrderStatusResp{}, nil
}

func (l *UpdateOrderStatusLogic) verifyOrderStatus(newStatus, oldStatus int64) error {
	if newStatus == globalKey.OrderWaitPay {
		return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_UPDATE_STATUS_ERROR),
			"无法将订单修改为未付款状态, newStatus: %v, oldStatus: %v", newStatus, oldStatus)
	}
	if newStatus == globalKey.OrderPayed {
		if oldStatus != globalKey.OrderWaitPay {
			return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_UPDATE_STATUS_ERROR),
				"未支付的订单无法修改成已付款的订单, newStatus: %v, oldStatus: %v", newStatus, oldStatus)
		}
	}
	if newStatus == globalKey.OrderCancel {
		if oldStatus != globalKey.OrderWaitPay {
			return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_UPDATE_STATUS_ERROR),
				"未支付的订单无法修改成已取消的订单, newStatus: %v, oldStatus: %v", newStatus, oldStatus)
		}
	}
	if newStatus == globalKey.OrderWaitShip {
		if oldStatus != globalKey.OrderPayed {
			return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_UPDATE_STATUS_ERROR),
				"非已支付的订单无法修改成待发货的订单, newStatus: %v, oldStatus: %v", newStatus, oldStatus)
		}
	}
	if newStatus == globalKey.OrderWaitReceive {
		if oldStatus != globalKey.OrderWaitShip {
			return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_UPDATE_STATUS_ERROR),
				"非待发货的订单无法修改成待收货的订单, newStatus: %v, oldStatus: %v", newStatus, oldStatus)
		}
	}
	if newStatus == globalKey.OrderSuccess {
		if oldStatus != globalKey.OrderWaitReceive {
			return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_UPDATE_STATUS_ERROR),
				"非待收货的订单无法修改成交易成功的订单, newStatus: %v, oldStatus: %v", newStatus, oldStatus)
		}
	}
	if newStatus == globalKey.OrderRefund {
		if oldStatus < globalKey.OrderPayed {
			return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_UPDATE_STATUS_ERROR),
				"未支付的订单无法修改成退款的订单, newStatus: %v, oldStatus: %v", newStatus, oldStatus)
		}
	}
	return nil
}
