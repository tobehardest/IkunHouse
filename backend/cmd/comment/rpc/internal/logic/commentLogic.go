package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"IkunHouse/cmd/comment/rpc/comment"
	"IkunHouse/cmd/comment/rpc/internal/svc"

	"IkunHouse/cmd/comment/code"
	"IkunHouse/cmd/comment/model"
	"IkunHouse/pkg/uniqueid/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentLogic) Comment(in *comment.CommentReq) (*comment.CommentRes, error) {
	// todo: add your logic here and delete this line
	_, err := snowflake.GetID()
	if err != nil {
		return nil, errors.Wrapf(code.ErrorGenIDFailed, "postId:%s", in.PostId)
	}

	commentData := &model.Comment{}
	err = copier.Copy(commentData, in)
	if err != nil {

	}
	//commentData.Id = commentID
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, commentData)
	if err != nil {

	}

	return &comment.CommentRes{}, nil
}
