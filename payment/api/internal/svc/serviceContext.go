package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/payment/api/internal/config"
	"zero-shop/payment/api/internal/middleware"
	"zero-shop/payment/rpc/payment"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config         config.Config
	CheckUserState rest.Middleware

	UserRpc    user.User
	PaymentRpc payment.Payment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CheckUserState: middleware.NewCheckUserStateMiddleware().Handle,

		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpc)),
	}
}
