package collect

import (
	"context"
	"github.com/jinzhu/copier"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"IkunHouse/cmd/collect/rpc/collect"
	"github.com/zeromicro/go-zero/core/logx"
)

type CollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectLogic {
	return &CollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectLogic) Collect(req *types.CollectReq) (resp *types.CollectRes, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CollectClient.Collect(l.ctx, &collect.CollectReq{})
	if err != nil {

	}
	err = copier.Copy(resp, res)
	if err != nil {

	}
	return resp, nil
}
