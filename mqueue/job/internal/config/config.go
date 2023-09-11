package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	RedisConf struct {
		Addr     string
		Password string
		DB       int
		PoolSize int
	}

	OrderRpc   zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
