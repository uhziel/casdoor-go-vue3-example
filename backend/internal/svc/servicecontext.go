package svc

import (
	"backend/internal/config"
	"backend/internal/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                config.Config
	AccessTokenMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(c.Cert))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:                c,
		AccessTokenMiddleware: middleware.NewAccessTokenMiddleware(publicKey).Handle,
	}
}
