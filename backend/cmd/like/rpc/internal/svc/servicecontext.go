package svc

import (
	"IkunHouse/cmd/like/model"
	"IkunHouse/cmd/like/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	LikeModel      model.LikeRecordModel
	LikeCountModel model.LikeCountModel
	BizRedis       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		LikeModel:      model.NewLikeRecordModel(conn, c.CacheRedis),
		LikeCountModel: model.NewLikeCountModel(conn, c.CacheRedis),
		BizRedis:       rds,
	}
}
