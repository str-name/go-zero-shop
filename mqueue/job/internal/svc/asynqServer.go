package svc

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-shop/mqueue/job/internal/config"
)

func newAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(asynq.RedisClientOpt{
		Addr:     c.RedisConf.Addr,
		Password: c.RedisConf.Password,
		DB:       c.RedisConf.DB,
		PoolSize: c.RedisConf.PoolSize,
	}, asynq.Config{
		IsFailure: func(err error) bool {
			logx.WithContext(context.Background()).Errorf("asynq server exec task IsFailure =====>>>>>  ERROR: %+v", err)
			return true
		},
		Concurrency: 20,
	})
}
