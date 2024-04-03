package svc

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                config.Config
	AccessTokenMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		AccessTokenMiddleware: middleware.NewAccessTokenMiddleware().Handle,
	}
}
