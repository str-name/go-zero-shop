package kq

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-shop/common/globalKey"
	"zero-shop/common/kqOrder"
	"zero-shop/common/xerr"
	"zero-shop/order/mq/internal/svc"
	"zero-shop/order/rpc/order"
)

type PaymentUpdateOrderStateMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateOrderStateMq(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentUpdateOrderStateMq {
	return &PaymentUpdateOrderStateMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentUpdateOrderStateMq) Consume(key, val string) error {

	// 获取队列信息
	var kqMessage kqOrder.PaymentUpdateOrderState
	if err := json.Unmarshal([]byte(val), &kqMessage); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateOrderStateMq`Consume Unmarshal kqOrder.PaymentUpdateOrderState ERROR: %+v", err)
		return err
	}

	// 判断订单是否存在
	orderResp, err := l.svcCtx.OrderRpc.GetOrderOnlyDetail(l.ctx, &order.GetOrderOnlyDetailReq{
		UserID:  kqMessage.UserID,
		OrderSn: kqMessage.OrderSn,
	})
	if err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateOrderStateMq`Consume USE OrderRpc.GetOrderOnlyDetail ERROR: %+v, userID: %v, orderSn: %v",
			err, kqMessage.UserID, kqMessage.OrderSn)
		return err
	}

	// 判断订单状态更新是否正确
	if orderResp.Status != globalKey.OrderWaitPay {
		logx.WithContext(l.ctx).Error("Order`State ERROR, userID: %v, orderSn: %v, oldState: %v, newState: %v",
			kqMessage.UserID, kqMessage.OrderSn, orderResp.Status, kqMessage.OrderState)
		return xerr.NewErrMsg("Order`State ERROR")
	}

	// 更新订单状态
	l.svcCtx.OrderRpc.UpdateOrderStatus(l.ctx, &order.UpdateOrderStatusReq{
		OrderSn:     kqMessage.OrderSn,
		UserID:      kqMessage.UserID,
		OrderStatus: kqMessage.OrderState,
	})

	return nil
}
