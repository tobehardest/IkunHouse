package main

import (
	"flag"
	"fmt"

	"video_clip/cmd/api/internal/config"
	"video_clip/cmd/api/internal/handler"
	"video_clip/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// local
// var configFile = flag.String("f", "etc/videoclip.yaml", "the config file")
// dev
var configFile = flag.String("f", "./cmd/api/etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
