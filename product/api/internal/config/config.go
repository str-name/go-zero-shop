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
	Qiniu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		CDN       string
		Zone      string
		Prefix    string
	}

	ProductRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
