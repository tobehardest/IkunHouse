package logic

import (
	"context"

	"IkunHouse/cmd/comment/rpc/comment"
	"IkunHouse/cmd/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *comment.DelCommentReq) (*comment.DelCommentRes, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.CommentModel.DelCommentById(l.ctx, in.CommentId)
	if err != nil {

	}

	return &comment.DelCommentRes{}, nil
}
