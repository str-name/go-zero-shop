package listen

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"zero-shop/order/mq/internal/config"
	kq2 "zero-shop/order/mq/internal/mqs/kq"
	"zero-shop/order/mq/internal/svc"
)

func KqMqs(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.PaymentUpdateOrderState, kq2.NewPaymentUpdateOrderStateMq(ctx, svcCtx)),
	}
}
