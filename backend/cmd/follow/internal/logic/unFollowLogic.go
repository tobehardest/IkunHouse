package logic

import (
	"context"

	"video_clip/cmd/follow/follow"
	"video_clip/cmd/follow/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消关注
func (l *UnFollowLogic) UnFollow(in *follow.UnFollowRequest) (*follow.UnFollowResponse, error) {
	// todo: add your logic here and delete this line

	return &follow.UnFollowResponse{}, nil
}
