package svc

import (
	"IkunHouse/cmd/video/model"
	"IkunHouse/cmd/video/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	SqlConn       sqlx.SqlConn
	VideoModel    model.TVideoModel
	TagModel      model.TagModel
	VideoTagModel model.VideoTagModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlConf.Datasource)
	return &ServiceContext{
		Config:        c,
		SqlConn:       conn,
		VideoModel:    model.NewTVideoModel(conn),
		TagModel:      model.NewTagModel(conn),
		VideoTagModel: model.NewVideoTagModel(conn),
	}
}
