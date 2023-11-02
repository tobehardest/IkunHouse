package comment

import (
	"context"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelCommentLogic {
	return &CancelCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelCommentLogic) CancelComment(req *types.CancelCommentReq) (resp *types.CancelCommentReq, err error) {
	// todo: add your logic here and delete this line

	return
}
