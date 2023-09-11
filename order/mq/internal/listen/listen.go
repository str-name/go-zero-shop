package listen

import (
	"context"
	"github.com/zeromicro/go-zero/core/service"
	"zero-shop/order/mq/internal/config"
	"zero-shop/order/mq/internal/svc"
)

func Mqs(c config.Config) []service.Service {
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	// 使用go-queue中的kq
	services = append(services, KqMqs(c, ctx, svcCtx)...)

	return services
}
