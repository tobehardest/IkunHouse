package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"video_clip/cmd/search/mq/internal/config"
	"video_clip/cmd/search/mq/internal/mqs"
	"video_clip/cmd/search/mq/internal/svc"
)

// var configFile = flag.String("f", "etc/mq.yaml", "the config file")

var configFile = flag.String("f", "./cmd/search/mq/etc/mq.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}
