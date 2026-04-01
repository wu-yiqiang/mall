package svc

import (
	"mall/service/user/api/internal/config"
	"mall/service/user/api/internal/middleware"
	"mall/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Cost      rest.Middleware
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn, c.CacheRedis),
		Cost:      middleware.NewCostMiddleware().Handle,
	}
}
