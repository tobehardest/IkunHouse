package collect

import (
	"context"
	"github.com/jinzhu/copier"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"IkunHouse/cmd/collect/rpc/collect"
	"github.com/zeromicro/go-zero/core/logx"
)

type CollectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectListLogic {
	return &CollectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectListLogic) CollectList(req *types.CollectListReq) (resp *types.CollectListRes, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CollectClient.CollectList(l.ctx, &collect.CollectListReq{})
	if err != nil {

	}
	err = copier.Copy(resp, res)
	if err != nil {

	}
	return resp, nil
}
