package logic

import (
	"context"
	"video_clip/cmd/video/rpc/internal/svc"
	"video_clip/cmd/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectLogic {
	return &AddCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCollectLogic) AddCollect(in *video.AddCollectRequest) (*video.AddCollectResponse, error) {
	// todo: add your logic here and delete this line

	return &video.AddCollectResponse{}, nil
}
