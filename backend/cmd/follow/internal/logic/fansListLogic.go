package logic

import (
	"context"

	"video_clip/cmd/follow/follow"
	"video_clip/cmd/follow/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FansListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansListLogic {
	return &FansListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 粉丝列表
func (l *FansListLogic) FansList(in *follow.FansListRequest) (*follow.FansListResponse, error) {
	// todo: add your logic here and delete this line

	return &follow.FansListResponse{}, nil
}
