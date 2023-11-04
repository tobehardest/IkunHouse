package main

import (
	"flag"
	"fmt"
	"video_clip/cmd/auth/rpc/authCenter"
	"video_clip/cmd/auth/rpc/internal/config"
	"video_clip/cmd/auth/rpc/internal/server"
	"video_clip/cmd/auth/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// prd
// var configFile = flag.String("f", "etc/auth.yaml", "the config file")
// dev
var configFile = flag.String("f", "./cmd/auth/rpc/etc/authCenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		authCenter.RegisterAuthCenterServer(grpcServer, server.NewAuthCenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
