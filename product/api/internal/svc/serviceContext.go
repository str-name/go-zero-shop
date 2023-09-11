package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/product/api/internal/config"
	"zero-shop/product/api/internal/middleware"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config          config.Config
	CheckStoreState rest.Middleware
	CheckUserState  rest.Middleware

	ProductRpc product.ProductZrpcClient
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CheckStoreState: middleware.NewCheckStoreStateMiddleware().Handle,
		CheckUserState:  middleware.NewCheckUserStateMiddleware().Handle,

		ProductRpc: product.NewProductZrpcClient(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
