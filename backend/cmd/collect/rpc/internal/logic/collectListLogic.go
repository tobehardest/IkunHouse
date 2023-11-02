package logic

import (
	"context"

	"video_clip/cmd/collect/rpc/collect"
	"video_clip/cmd/collect/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectListLogic {
	return &CollectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏列表
func (l *CollectListLogic) CollectList(in *collect.CollectListReq) (*collect.CollectListRes, error) {
	// todo: add your logic here and delete this line

	return &collect.CollectListRes{}, nil
}
