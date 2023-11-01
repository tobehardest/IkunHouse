package logic

import (
	"context"

	"video_clip/cmd/video/internal/svc"
	"video_clip/cmd/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadVideoLogic) UploadVideo(in *video.UploadVideoReq) (*video.UploadVideoRes, error) {
	// todo: add your logic here and delete this line

	return &video.UploadVideoRes{}, nil
}
