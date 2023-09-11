package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/order/api/internal/config"
	"zero-shop/order/api/internal/middleware"
	"zero-shop/order/rpc/order"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config         config.Config
	CheckUserState rest.Middleware

	OrderRpc order.Order
	UserRpc  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CheckUserState: middleware.NewCheckUserStateMiddleware().Handle,

		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
