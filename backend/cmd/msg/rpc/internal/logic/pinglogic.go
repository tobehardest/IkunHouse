package logic

import (
	"IkunHouse/cmd/msg/rpc/internal/svc"
	"IkunHouse/cmd/msg/rpc/msg"
	"context"

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
