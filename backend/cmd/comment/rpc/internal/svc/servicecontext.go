package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"video_clip/cmd/comment/model"
	"video_clip/cmd/comment/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(conn),
	}
}
