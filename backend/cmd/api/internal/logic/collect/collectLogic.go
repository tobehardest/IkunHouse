package collect

import (
	"context"
	"github.com/jinzhu/copier"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/collect/rpc/collect"
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
