package main

import (
	"flag"
	"fmt"
	"video_clip/cmd/api-service/internal/config"
	"video_clip/cmd/api-service/internal/handler"
	"video_clip/cmd/api-service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//httpx.SetErrorHandler(func(err error) (int,interface{}) {
	//	switch e:=err.(type) {
	//	case apiErr.:
	//
	//	}
	//}

	//logx.DisableStat()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
