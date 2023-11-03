package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"video_clip/cmd/auth/model"
	"video_clip/cmd/auth/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		//UserModel: model.NewUserModel(conn, c.CacheRedis),
		UserModel: model.NewUserModel(conn),
	}
}
