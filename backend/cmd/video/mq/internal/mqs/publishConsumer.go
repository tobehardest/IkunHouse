package mqs

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/video/mq/internal/svc"
)

type PublishConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func (l *PublishConsumer) Consume(key, val string) error {
	fmt.Printf("get key: %s val: %s\n", key, val)
	return nil
}

func NewPublishConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *PublishConsumer {
	return &PublishConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
