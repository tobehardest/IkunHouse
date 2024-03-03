package video

import (
	"IkunHouse/cmd/video/rpc/video"
	"context"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTypeListLogic {
	return &GetTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTypeListLogic) GetTypeList(req *types.GetTypeListRequest) (resp *types.GetTypeListResponse, err error) {
	in := &video.GetTypeListRequest{}
	type_list, err := l.svcCtx.VideoClient.GetTypeList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	t_list := make([]types.Type, len(type_list.TypeList))
	for index, t := range type_list.TypeList {
		t_list[index].Id = t.Id
		t_list[index].Name = t.Name
	}
	resp.TypeList = t_list
	return
}
