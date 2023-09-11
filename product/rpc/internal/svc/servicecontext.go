package svc

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"zero-shop/common/goredis"
	gorm2 "zero-shop/common/gorm"
	"zero-shop/product/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	ProductDB *gorm.DB
	RedisDB   *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		ProductDB: gorm2.ProductDB,
		RedisDB:   goredis.Rdb,
	}
}
