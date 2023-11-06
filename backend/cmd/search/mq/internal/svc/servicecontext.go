package svc

import (
	"video_clip/cmd/search/mq/internal/config"
	"video_clip/pkg/es"
)

type ServiceContext struct {
	Config config.Config
	Es     *es.Es
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Es: es.MustNewEs(&es.Config{
			Addresses: c.Es.Addresses,
			Username:  c.Es.Username,
			Password:  c.Es.Password,
		}),
	}
}
