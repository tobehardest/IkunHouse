package logic

import (
	"context"

	"video_clip/cmd/follow/follow"
	"video_clip/cmd/follow/internal/svc"

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

	return &follow.FollowResponse{}, nil
}
