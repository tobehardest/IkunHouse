package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"video_clip/cmd/follow/model"
	"video_clip/cmd/follow/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	FollowModel      *model.FollowModel
	FollowCountModel *model.FollowCountModel
	BizRedis         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:           c,
		FollowModel:      model.NewFollowModel(db.DB),
		FollowCountModel: model.NewFollowCountModel(db.DB),
		BizRedis:         rds,
	}
}
