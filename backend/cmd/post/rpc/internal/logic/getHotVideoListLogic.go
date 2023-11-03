package logic

import (
	"context"

	"video_clip/cmd/post/rpc/internal/svc"
	"video_clip/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHotVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotVideoListLogic {
	return &GetHotVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获得热门视频
func (l *GetHotVideoListLogic) GetHotVideoList(in *post.GetHotVideoRequest) (*post.GetHotVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &post.GetHotVideoResponse{}, nil
}
