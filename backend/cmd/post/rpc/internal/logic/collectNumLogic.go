package logic

import (
	"context"

	"video_clip/cmd/post/rpc/internal/svc"
	"video_clip/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectNumLogic {
	return &CollectNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏数操作
func (l *CollectNumLogic) CollectNum(in *post.CollectNumRequest) (*post.CollectNumResponse, error) {
	if in.Action == 1 {
		in.Count = -1 * in.Count
	}
	err := l.svcCtx.PostModel.UpdateCollectNum(l.ctx, in.Vid, in.Count)
	return &post.CollectNumResponse{}, err
}
