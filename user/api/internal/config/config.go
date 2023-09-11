package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	MysqlConf struct {
		DSN string
	}
	RedisConf struct {
		Addr     string
		Password string
		DB       int
		PoolSize int
	}
	Email struct {
		Host     string
		Port     int64
		User     string
		Password string
	}
	Qiniu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		CDN       string
		Zone      string
		Prefix    string
	}

	UserRpc zrpc.RpcClientConf
}
