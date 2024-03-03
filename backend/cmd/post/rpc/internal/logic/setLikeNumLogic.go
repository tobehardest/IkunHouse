package logic

import (
	"context"

	"IkunHouse/cmd/post/rpc/internal/svc"
	"IkunHouse/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLikeNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetLikeNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLikeNumLogic {
	return &SetLikeNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞数操作
func (l *SetLikeNumLogic) SetLikeNum(in *post.LikeNumRequest) (*post.LikeNumResponse, error) {
	if in.Action == 1 {
		in.Count = -1 * in.Count
	}
	err := l.svcCtx.PostModel.UpdateLikeNum(l.ctx, in.Vid, in.Count)

	return &post.LikeNumResponse{}, err
}
