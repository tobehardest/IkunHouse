package svc

import (
	"IkunHouse/cmd/comment/model"
	"IkunHouse/cmd/comment/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
