// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"api/internal/config"
	"mall/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
