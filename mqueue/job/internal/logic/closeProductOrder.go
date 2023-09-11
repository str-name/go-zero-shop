package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/mqueue/job/internal/svc"
	"zero-shop/mqueue/job/jobtype"
	"zero-shop/order/rpc/order"
)

var ErrCloseOrderFail = xerr.NewErrMsg("CLOSE ORDER FAIL")

type CloseProductOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseProductOrderHandler(svcCtx *svc.ServiceContext) *CloseProductOrderHandler {
	return &CloseProductOrderHandler{svcCtx: svcCtx}
}

func (l *CloseProductOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var payload jobtype.DeferCloseSeckillOrderPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return errors.Wrapf(ErrCloseOrderFail, "CloseProductOrderHandler Fail, payload: %+v, ERROR: %+v", t.Payload(), err)
	}

	orderResp, err := l.svcCtx.OrderRpc.OrderDetail(ctx, &order.OrderDetailReq{
		UserID:  payload.UserID,
		OrderSn: payload.OrderSn,
	})
	if err != nil || orderResp == nil {
		return errors.Wrapf(ErrCloseOrderFail, "CloseProductOrderHandler GET order Fail OR order no exists, OrderSn: %v, ERROR: %+v", payload.OrderSn, err)
	}

	if orderResp.Status == globalKey.OrderWaitPay {
		_, err := l.svcCtx.OrderRpc.UpdateOrderStatus(ctx, &order.UpdateOrderStatusReq{
			OrderSn:     payload.OrderSn,
			UserID:      payload.UserID,
			OrderStatus: globalKey.OrderCancel,
		})
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFail, "CloseProductOrderHandler CLOSE order Fail, OrderSn: %v, ERROR: %+v", payload.OrderSn, err)
		}
	}

	return nil
}
