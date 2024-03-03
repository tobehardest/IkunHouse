package mqs

import (
	"IkunHouse/cmd/search/mq/internal/config"
	"IkunHouse/cmd/search/mq/internal/svc"
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.ShortVideoKqConsumerConf, NewPublishConsumer(ctx, svcContext)),
		//kq.MustNewQueue(c.UserKqConsumerConf, NewPublishConsumer(ctx, svcContext)),
		//.....
	}

}
