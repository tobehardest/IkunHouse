package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"video_clip/service/api-service/internal/config"
	"video_clip/service/api-service/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,
	}
}