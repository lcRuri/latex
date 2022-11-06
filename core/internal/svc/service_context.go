package svc

import (
	"core/internal/config"
	"core/internal/middleware"
	"core/models"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DataSource),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
