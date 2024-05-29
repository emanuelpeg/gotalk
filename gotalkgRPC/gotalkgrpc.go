package main

import (
	"flag"
	"fmt"
	"gotalk/gotalk"
	"gotalk/internal/config"
	"gotalk/internal/server"
	"gotalk/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/gotalkgrpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		gotalk.RegisterHelloServer(grpcServer, server.NewHelloServer(ctx))
		gotalk.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(ctx))
		gotalk.RegisterHealthCheckServer(grpcServer, server.NewHealthCheckServer(ctx))
		gotalk.RegisterChatServiceServer(grpcServer, server.NewChatServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
