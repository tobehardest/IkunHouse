package logic

import (
	"context"

	"video_clip/cmd/video/rpc/internal/svc"
	"video_clip/cmd/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectByVidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectByVidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectByVidLogic {
	return &GetCollectByVidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCollectByVidLogic) GetCollectByVid(in *video.GetCollectByVidRequest) (*video.GetCollectByVidResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetCollectByVidResponse{}, nil
}
