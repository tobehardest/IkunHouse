package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"IkunHouse/cmd/comment/rpc/comment"
	"IkunHouse/cmd/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.CommentListReq) (*comment.CommentListRes, error) {
	// todo: add your logic here and delete this line
	ret, err := l.svcCtx.CommentModel.GetCommentListByIDs(l.ctx, in.Ids)
	if err != nil {

	}
	out := &comment.CommentListRes{}
	err = copier.Copy(out, ret)
	if err != nil {

	}

	return &comment.CommentListRes{}, nil
}
