package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/service/system/api/internal/config"
	"mall/service/system/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	RoleModel model.RoleModel
	DB        sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn, c.CacheRedis),
		RoleModel: model.NewRoleModel(sqlxConn, c.CacheRedis),
		DB:        sqlxConn,
	}
}
