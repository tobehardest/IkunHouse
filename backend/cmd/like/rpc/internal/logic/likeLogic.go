package logic

import (
	"context"

	"video_clip/cmd/like/rpc/internal/svc"
	"video_clip/cmd/like/rpc/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeLogic) Like(in *like.LikeReq) (*like.LikeRes, error) {
	// todo: add your logic here and delete this line

	return &like.LikeRes{}, nil
}
