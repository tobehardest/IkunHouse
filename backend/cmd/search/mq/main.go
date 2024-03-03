package main

import (
	"IkunHouse/cmd/search/mq/internal/config"
	"IkunHouse/cmd/search/mq/internal/mqs"
	"IkunHouse/cmd/search/mq/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
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
