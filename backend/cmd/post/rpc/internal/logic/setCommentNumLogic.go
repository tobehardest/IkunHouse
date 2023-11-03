package logic

import (
	"context"

	"video_clip/cmd/post/rpc/internal/svc"
	"video_clip/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCommentNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetCommentNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCommentNumLogic {
	return &SetCommentNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论数操作
func (l *SetCommentNumLogic) SetCommentNum(in *post.CommentNumRequest) (*post.CommentNumResponse, error) {
	if in.Action == 1 {
		in.Count = -1 * in.Count
	}
	err := l.svcCtx.PostModel.UpdateCommentNum(l.ctx, in.Vid, in.Count)

	return &post.CommentNumResponse{}, err
}
