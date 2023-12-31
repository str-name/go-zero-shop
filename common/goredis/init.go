package goredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

var Rdb *redis.Client

func init() {
	option := redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       4,
		PoolSize: 100,
	}
	rdb := redis.NewClient(&option)
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logx.WithContext(context.Background()).Error("Redis connect ERROR: %+v", err)
	}
	Rdb = rdb
}
