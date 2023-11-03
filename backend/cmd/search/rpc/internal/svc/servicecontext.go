package svc

import (
	"video_clip/cmd/search/rpc/internal/config"
	"video_clip/pkg/es"
)

type ServiceContext struct {
	Config config.Config
	Es     *es.TypedEs
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Es: es.MustNewTypedEs(&es.Config{
			Addresses: c.Es.Addresses,
			Username:  c.Es.Username,
			Password:  c.Es.Password,
		}),
	}
}
