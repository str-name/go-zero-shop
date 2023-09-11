package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	MysqlConf struct {
		DSN string
	}
	RedisConf struct {
		Addr     string
		Password string
		DB       int
		PoolSize int
	}

	Qiniu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		CDN       string
		Zone      string
		Prefix    string
	}
}
