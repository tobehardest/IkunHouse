package logic

import (
	"IkunHouse/cmd/video/rpc/video"
	"context"

	"IkunHouse/cmd/video/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReadCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddReadCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReadCountLogic {
	return &AddReadCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddReadCountLogic) AddReadCount(in *video.AddReadCountRequest) (*video.AddReadCountResponse, error) {
	err := l.svcCtx.VideoModel.AddReadCount(context.Background(), in.Id, in.Count)
	return &video.AddReadCountResponse{}, err
}
