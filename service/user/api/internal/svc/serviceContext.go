package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/service/user/api/internal/config"
	"mall/service/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn),
	}
}
