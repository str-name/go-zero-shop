package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"zero-shop/common/goredis"
	gorm2 "zero-shop/common/gorm"
	"zero-shop/user/api/internal/config"
	"zero-shop/user/api/internal/middleware"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config         config.Config
	CheckUserState rest.Middleware
	UserDB         *gorm.DB
	RedisDB        *redis.Client
	UserRpc        user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CheckUserState: middleware.NewCheckUserStateMiddleware().Handle,

		UserDB:  gorm2.UserDB,
		RedisDB: goredis.Rdb,

		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
