package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"video_clip/cmd/collect/model"
	"video_clip/cmd/collect/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	CollectModel      model.CollectRecordModel
	CollectCountModel model.CollectCountModel
	BizRedis          *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)

	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	return &ServiceContext{
		Config:            c,
		CollectModel:      model.NewCollectRecordModel(conn),
		CollectCountModel: model.NewCollectCountModel(conn),
		BizRedis:          rds,
	}
}
