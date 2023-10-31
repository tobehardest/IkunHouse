package logic

import (
	"context"

	"video_clip/cmd/video/internal/svc"
	"video_clip/cmd/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *video.Request) (*video.Response, error) {
	// todo: add your logic here and delete this line

	return &video.Response{}, nil
}
