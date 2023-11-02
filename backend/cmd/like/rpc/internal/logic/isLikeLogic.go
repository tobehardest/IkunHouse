package logic

import (
	"context"

	"video_clip/cmd/like/rpc/internal/svc"
	"video_clip/cmd/like/rpc/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLikeLogic {
	return &IsLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsLikeLogic) IsLike(in *like.IsLikeReq) (*like.IsLikeRes, error) {
	// todo: add your logic here and delete this line

	return &like.IsLikeRes{}, nil
}
