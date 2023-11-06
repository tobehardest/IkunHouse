package logic

import (
	"context"

	"video_clip/cmd/like/rpc/internal/svc"
	"video_clip/cmd/like/rpc/like"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/pkg/utils"
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
	// 1. 查询redis
	likeKey := utils.GetEntityLikeKey(in.BizId, in.ObjId)
	_, err := l.svcCtx.BizRedis.SismemberCtx(l.ctx, likeKey, in.UserId)
	if err != nil {

	}
	out := &like.IsLikeRes{}
	return out, nil
}
