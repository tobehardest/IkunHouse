package svc

import (
	"IkunHouse/cmd/search/mq/internal/config"
	"IkunHouse/pkg/es"
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
