package logic

import (
	"context"

	"IkunHouse/cmd/post/rpc/internal/svc"
	"IkunHouse/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoByUidLogic {
	return &GetVideoByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据用户id查找视频
func (l *GetVideoByUidLogic) GetVideoByUid(in *post.GetVideoByUidRequest) (*post.GetVideoByUidResponse, error) {
	// todo: add your logic here and delete this line

	return &post.GetVideoByUidResponse{}, nil
}
