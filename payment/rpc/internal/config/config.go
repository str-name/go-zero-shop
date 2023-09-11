package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	UserRpc  zrpc.RpcClientConf
	OrderRpc zrpc.RpcClientConf

	KqPaymentUpdateOrderStateConf struct {
		Brokers []string
		Topic   string
	}
}
