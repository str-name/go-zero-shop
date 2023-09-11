package svc

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"zero-shop/common/goredis"
	gorm2 "zero-shop/common/gorm"
	"zero-shop/user/rpc/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserDB  *gorm.DB
	RedisDB *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserDB:  gorm2.UserDB,
		RedisDB: goredis.Rdb,
	}
}
