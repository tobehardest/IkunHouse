package follow

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"IkunHouse/cmd/follow/rpc/follow"
	"IkunHouse/pkg/errx"
	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowRes, err error) {
	out, err := l.svcCtx.FollowClient.Follow(l.ctx, &follow.FollowReq{
		UserId:         req.UserId,
		FollowedUserId: req.FollowedUserId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.FollowRes{}
	err = copier.Copy(resp, out)
	if err != nil {
		errors.Wrapf(errx.NewErrCode(errx.COPIER_COPY_ERROR), "userId:%s", req.UserId)
	}
	return
}
