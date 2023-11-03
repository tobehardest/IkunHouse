package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"video_clip/cmd/comment/rpc/comment"
	"video_clip/cmd/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"video_clip/cmd/comment/code"
	"video_clip/cmd/comment/model"
	"video_clip/pkg/uniqueid/snowflake"
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
