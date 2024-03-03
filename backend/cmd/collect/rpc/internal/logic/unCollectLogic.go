package logic

import (
	"context"

	"IkunHouse/cmd/collect/rpc/collect"
	"IkunHouse/cmd/collect/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnCollectLogic {
	return &UnCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消收藏
func (l *UnCollectLogic) UnCollect(in *collect.UnCollectReq) (*collect.UnCollectRes, error) {
	// todo: add your logic here and delete this line

	return &collect.UnCollectRes{}, nil
}
