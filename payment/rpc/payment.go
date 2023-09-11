package main

import (
	"flag"
	"fmt"

	"zero-shop/common/intercepter/rpcLoggerIntercepter"

	"zero-shop/payment/rpc/internal/config"
	"zero-shop/payment/rpc/internal/server"
	"zero-shop/payment/rpc/internal/svc"
	"zero-shop/payment/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/payment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPaymentServer(grpcServer, server.NewPaymentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	s.AddUnaryInterceptors(rpcLoggerIntercepter.RpcLoggerIntercepter)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
