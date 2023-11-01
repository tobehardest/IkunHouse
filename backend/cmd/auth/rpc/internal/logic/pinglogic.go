package logic

import (
	"context"
	"video_clip/cmd/auth/rpc/auth"
	"video_clip/cmd/auth/rpc/internal/svc"

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

func (l *PingLogic) Ping(in *auth.Request) (*auth.Response, error) {
	// todo: add your logic here and delete this line

	return &auth.Response{}, nil
}
