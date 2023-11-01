package logic

import (
	"context"
	"video_clip/cmd/video/rpc/video"

	"video_clip/cmd/video/rpc/internal/svc"

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
	// todo: add your logic here and delete this line

	return &video.AddReadCountResponse{}, nil
}
