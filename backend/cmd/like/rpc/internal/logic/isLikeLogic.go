package logic

import (
	"context"

	"IkunHouse/cmd/like/rpc/internal/svc"
	"IkunHouse/cmd/like/rpc/like"

	"IkunHouse/pkg/utils"
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
	// 1. 查询redis
	likeKey := utils.GetEntityLikeKey(in.BizId, in.ObjId)
	_, err := l.svcCtx.BizRedis.SismemberCtx(l.ctx, likeKey, in.UserId)
	if err != nil {

	}
	out := &like.IsLikeRes{}
	return out, nil
}
