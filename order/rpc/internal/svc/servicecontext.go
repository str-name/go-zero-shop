package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"zero-shop/common/goredis"
	gorm2 "zero-shop/common/gorm"
	"zero-shop/order/rpc/internal/config"
	"zero-shop/payment/rpc/payment"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	OrderDB *gorm.DB
	RedisDB *redis.Client

	ProductRpc product.ProductZrpcClient
	UserRpc    user.User
	PaymentRpc payment.Payment

	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderDB: gorm2.OrderDB,
		RedisDB: goredis.Rdb,

		ProductRpc: product.NewProductZrpcClient(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpc)),

		AsynqClient: newAsynqClient(c),
	}
}
