package logic

import (
	"IkunHouse/cmd/video/model"
	"context"

	"IkunHouse/cmd/video/rpc/internal/svc"
	"IkunHouse/cmd/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfExistVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIfExistVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfExistVideoLogic {
	return &IfExistVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IfExistVideoLogic) IfExistVideo(in *video.ExistVideoRequest) (*video.ExistVideoResponse, error) {
	v, err := l.svcCtx.VideoModel.SelectVideoExist(l.ctx, in.VideoHash256)
	if err != nil {
		if err == model.ErrNotFound {
			return &video.ExistVideoResponse{Exist: false}, nil
		}
		return &video.ExistVideoResponse{}, err
	}

	return &video.ExistVideoResponse{Exist: true, Media: v.Media.String}, nil
}
