package logic

import (
	"context"
	"video_clip/cmd/post/rpc/internal/svc"
	"video_clip/cmd/post/rpc/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获得类别
func (l *GetTagListLogic) GetTagList(in *post.GetTagListRequest) (*post.GetTagListResponse, error) {
	tags, err := l.svcCtx.TagModel.FindAllTag(l.ctx)
	type_list := make([]*post.Tag, len(tags))
	for index, tag := range tags {
		type_list[index].Id = tag.Id
		type_list[index].Name = tag.TagName
	}
	return &post.GetTagListResponse{
		TagList: type_list,
	}, err
}
