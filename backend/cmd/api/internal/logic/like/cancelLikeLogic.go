package like

import (
	"context"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeLogic {
	return &CancelLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLikeLogic) CancelLike(req *types.CancelLikeReq) (resp *types.CancelLikeReq, err error) {
	// todo: add your logic here and delete this line

	return
}
