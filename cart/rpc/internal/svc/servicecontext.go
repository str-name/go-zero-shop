package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"zero-shop/cart/rpc/internal/config"
	"zero-shop/common/goredis"
	gorm2 "zero-shop/common/gorm"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	CartDB  *gorm.DB
	RedisDB *redis.Client

	ProductRpc product.ProductZrpcClient
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		CartDB:  gorm2.CartDB,
		RedisDB: goredis.Rdb,

		ProductRpc: product.NewProductZrpcClient(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
