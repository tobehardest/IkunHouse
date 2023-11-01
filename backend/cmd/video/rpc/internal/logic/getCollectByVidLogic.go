package logic

import (
	"context"
	"video_clip/cmd/video/rpc/internal/svc"
	"video_clip/cmd/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectByVidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectByVidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectByVidLogic {
	return &GetCollectByVidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCollectByVidLogic) GetCollectByVid(in *video.GetCollectByVidRequest) (*video.GetCollectByVidResponse, error) {
	//conn := model.NewVcVideoCollectModel(l.svcCtx.SqlConn)
	//conn.FindOne()

	return &video.GetCollectByVidResponse{}, nil
}
