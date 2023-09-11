package middleware

import "net/http"
import (
    "gorm.io/gorm"
    "zero-shop/common/goredis"
    gorm2 "zero-shop/common/gorm"
    "github.com/go-redis/redis/v8"
)

type {{.name}} struct {
    mysqlDB *gorm.DB
    redisDB *redis.Client
}

func New{{.name}}() *{{.name}} {
	return &{{.name}}{
	    mysqlDB: gorm2.UserDB,
	    redisDB: goredis.Rdb,
	}
}

func (m *{{.name}})Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
