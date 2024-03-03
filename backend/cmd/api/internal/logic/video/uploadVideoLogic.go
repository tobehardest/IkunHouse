package video

import (
	"IkunHouse/cmd/video/rpc/video"
	"context"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVideoLogic) UploadVideo(req *types.UploadVideoReq) (resp *types.UploadVideoRes, err error) {
	// 调用rpc服务进行上传
	in := &video.UploadVideoRequest{
		Media:       req.Media,
		Uid:         req.Uid,
		Title:       req.Title,
		CoverUrl:    req.CoverUrl,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Address:     req.Address,
		TagId:       req.TagId,
		VideoSha256: req.Sha256,
	}
	_, err = l.svcCtx.VideoClient.UploadVideo(l.ctx, in)
	if err != nil {
		return &types.UploadVideoRes{}, err
	}
	return &types.UploadVideoRes{}, nil
}
