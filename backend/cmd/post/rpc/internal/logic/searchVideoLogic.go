package logic

import (
	"context"

	"video_clip/cmd/post/rpc/internal/svc"
	"video_clip/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchVideoLogic {
	return &SearchVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据搜索查询视频
func (l *SearchVideoLogic) SearchVideo(in *post.SearchVideoRequest) (*post.SearchVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &post.SearchVideoResponse{}, nil
}
