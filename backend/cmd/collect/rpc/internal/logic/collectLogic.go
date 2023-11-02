package logic

import (
	"context"

	"video_clip/cmd/collect/rpc/collect"
	"video_clip/cmd/collect/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectLogic {
	return &CollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏
func (l *CollectLogic) Collect(in *collect.CollectReq) (*collect.CollectRes, error) {
	// todo: add your logic here and delete this line

	return &collect.CollectRes{}, nil
}
