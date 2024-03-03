package follow

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/follow/rpc/follow"
	"video_clip/pkg/errx"
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
