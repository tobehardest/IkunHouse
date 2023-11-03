package mqs

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"video_clip/cmd/search/mq/internal/config"
	"video_clip/cmd/search/mq/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.ShortVideoKqConsumerConf, NewPublishConsumer(ctx, svcContext)),
		//kq.MustNewQueue(c.UserKqConsumerConf, NewPublishConsumer(ctx, svcContext)),
		//.....
	}

}
