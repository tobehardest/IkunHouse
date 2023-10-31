package logic

import (
	"context"

	"video_clip/cmd/msg/internal/svc"
	"video_clip/cmd/msg/msg"

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

func (l *PingLogic) Ping(in *msg.Request) (*msg.Response, error) {
	// todo: add your logic here and delete this line

	return &msg.Response{}, nil
}
