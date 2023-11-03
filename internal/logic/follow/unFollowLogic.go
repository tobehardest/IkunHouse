package follow

import (
	"context"

	"Qiniu/internal/svc"
	"Qiniu/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnFollowLogic) UnFollow(req *types.UnFollowReq) (resp *types.UnFollowRes, err error) {
	// todo: add your logic here and delete this line

	return
}
