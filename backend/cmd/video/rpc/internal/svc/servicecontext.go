package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"video_clip/cmd/video/rpc/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	SqlConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlConf.Datasource)
	return &ServiceContext{
		Config:  c,
		SqlConn: conn,
	}
}
