package logic

import (
	"context"

	"video_clip/cmd/like/rpc/internal/svc"
	"video_clip/cmd/like/rpc/like"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/pkg/utils"
)

type LikeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeCountLogic {
	return &LikeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeCountLogic) LikeCount(in *like.LikeCountReq) (*like.LikeCountRes, error) {
	// todo: add your logic here and delete this line
	likeKey := utils.GetEntityLikeKey(in.BizId, in.ObjId)
	scardRet, err := l.svcCtx.BizRedis.ScardCtx(l.ctx, likeKey)
	if err != nil {

	}
	out := &like.LikeCountRes{
		LikeNum: scardRet,
	}
	return out, nil
}
