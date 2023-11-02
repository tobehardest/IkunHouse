package video

import (
	"context"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListByCatagoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoListByCatagoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListByCatagoryLogic {
	return &GetVideoListByCatagoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoListByCatagoryLogic) GetVideoListByCatagory(req *types.GetVideoListByCatagoryReq) (resp *types.GetVideoListByCatagoryRes, err error) {
	// todo: add your logic here and delete this line

	return
}
