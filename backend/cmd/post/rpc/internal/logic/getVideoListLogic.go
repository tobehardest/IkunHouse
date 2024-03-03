package logic

import (
	"context"

	"video_clip/cmd/post/rpc/internal/svc"
	"video_clip/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获得视频 - 默认/分类
func (l *GetVideoListLogic) GetVideoList(in *post.GetVideoListRequest) (*post.GetVideoListResponse, error) {
	// todo: add your logic here and delete this line

	return &post.GetVideoListResponse{}, nil
}
