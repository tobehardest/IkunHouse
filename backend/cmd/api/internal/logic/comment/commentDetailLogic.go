package comment

import (
	"context"

	"video_clip/cmd/api/internal/svc"
	"video_clip/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentDetailLogic {
	return &CommentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentDetailLogic) CommentDetail(req *types.CommentDetailReq) (resp *types.CommentDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
