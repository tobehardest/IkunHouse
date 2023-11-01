package logic

import (
	"context"
	"video_clip/cmd/follow/rpc/follow"
	"video_clip/cmd/follow/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注列表
func (l *FollowListLogic) FollowList(in *follow.FollowListRequest) (*follow.FollowListResponse, error) {
	// todo: add your logic here and delete this line

	return &follow.FollowListResponse{}, nil
}
