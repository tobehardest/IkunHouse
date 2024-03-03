package video

import (
	"context"

	"IkunHouse/cmd/api/internal/svc"
	"IkunHouse/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotVideoListLogic {
	return &GetHotVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotVideoListLogic) GetHotVideoList(req *types.GetHotVideoListReq) (resp *types.UploadVideoRes, err error) {
	// todo: add your logic here and delete this line

	return
}
