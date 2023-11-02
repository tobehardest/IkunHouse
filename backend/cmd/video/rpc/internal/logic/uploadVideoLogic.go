package logic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/video/model"
	"video_clip/cmd/video/rpc/internal/svc"
	"video_clip/cmd/video/rpc/video"
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

func (l *UploadVideoLogic) UploadVideo(in *video.UploadVideoRequest) (*video.UploadVideoResponse, error) {
	v := model.TVideo{
		Uid:      in.Uid,
		Title:    in.Title,
		Media:    sql.NullString{String: in.Media},
		CoverUrl: sql.NullString{String: in.CoverUrl},
		Type:     2,
		//Address: sql.NullString{Add}
		Longitude: in.Longitude,
		Latitude:  in.Latitude,
	}
	//result, err := conn.Insert(l.ctx, &v)
	result, err := l.svcCtx.VideoModel.Insert(l.ctx, &v)
	if err != nil {
		return &video.UploadVideoResponse{}, err
	}
	id, err := result.LastInsertId()
	return &video.UploadVideoResponse{
		Id: id,
	}, err
}
