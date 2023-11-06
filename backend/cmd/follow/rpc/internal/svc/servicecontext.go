package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"video_clip/cmd/follow/model"
	"video_clip/cmd/follow/rpc/internal/config"
	"video_clip/pkg/orm"
)

type ServiceContext struct {
	Config           config.Config
	DB               *orm.DB
	FollowModel      *model.FollowModel
	FollowCountModel *model.FollowCountModel
	BizRedis         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})

	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	return &ServiceContext{
		Config:           c,
		FollowModel:      model.NewFollowModel(db.DB),
		FollowCountModel: model.NewFollowCountModel(db.DB),
		BizRedis:         rds,
	}
}
