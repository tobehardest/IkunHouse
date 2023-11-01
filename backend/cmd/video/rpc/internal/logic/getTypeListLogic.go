package logic

import (
	"context"

	"video_clip/cmd/video/rpc/internal/svc"
	"video_clip/cmd/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeListLogic {
	return &GetTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTypeListLogic) GetTypeList(in *video.GetTypeListRequest) (*video.GetTypeListResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetTypeListResponse{}, nil
}
