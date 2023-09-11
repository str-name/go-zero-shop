package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-shop/common/goredis"
	"zero-shop/mqueue/job/internal/config"
	"zero-shop/order/rpc/order"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server

	RedisDB *redis.Client

	OrderRpc   order.Order
	ProductRpc product.ProductZrpcClient
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AsynqServer: newAsynqServer(c),

		RedisDB: goredis.Rdb,

		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: product.NewProductZrpcClient(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
