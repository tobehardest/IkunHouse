package logic

import (
	"context"
	"video_clip/cmd/follow/rpc/follow"
	"video_clip/cmd/follow/rpc/internal/svc"
	"video_clip/pkg/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注
func (l *FollowLogic) Follow(in *follow.FollowRequest) (*follow.FollowResponse, error) {
	// todo: add your logic here and delete this line
	if in.UserId == 0 {
		return nil, errx.NewErrCode(errx.FollowUserIdEmpty)
	}
	if in.FollowedUserId == 0 {
		return nil, errx.NewErrCode(errx.FollowUserIdEmpty)
	}
	if in.UserId == in.FollowedUserId {
		return nil, errx.NewErrCode(errx.CannotFollowSelf)
	}

	return &follow.FollowResponse{}, nil
}
