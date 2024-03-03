package svc

import (
	"IkunHouse/cmd/auth/model"
	"IkunHouse/cmd/auth/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
