package logic

import (
	"context"

	"video_clip/cmd/video/rpc/internal/svc"
	"video_clip/cmd/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectByUidLogic {
	return &GetCollectByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCollectByUidLogic) GetCollectByUid(in *video.GetCollectByUidRequest) (*video.GetCollectByUidResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetCollectByUidResponse{}, nil
}
