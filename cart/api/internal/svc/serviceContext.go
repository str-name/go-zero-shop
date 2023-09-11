package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/cart/api/internal/config"
	"zero-shop/cart/api/internal/middleware"
	"zero-shop/cart/rpc/cart"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config         config.Config
	CheckUserState rest.Middleware

	CartRpc cart.Cart
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CheckUserState: middleware.NewCheckUserStateMiddleware().Handle,
		
		CartRpc: cart.NewCart(zrpc.MustNewClient(c.CartRpc)),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
