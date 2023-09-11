package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	ProductRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	PaymentRpc zrpc.RpcClientConf
}
