package logic

import (
	"context"
	"video_clip/cmd/video/rpc/internal/svc"
	"video_clip/cmd/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeListLogic {
	return &GetTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTypeListLogic) GetTypeList(in *video.GetTypeListRequest) (*video.GetTypeListResponse, error) {
	tags, err := l.svcCtx.TagModel.FindAllTag(l.ctx)
	type_list := make([]*video.Type, len(tags))
	for index, tag := range tags {
		type_list[index].Id = tag.Id
		type_list[index].Name = tag.TagName
	}
	return &video.GetTypeListResponse{
		TypeList: type_list,
	}, err
}
