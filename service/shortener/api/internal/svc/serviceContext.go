// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"mall/service/shortener/api/internal/config"
	"mall/service/shortener/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	ShortUrlMapModel model.ShortUrlMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	return &ServiceContext{
		Config:           c,
		ShortUrlMapModel: model.NewShortUrlMapModel(sqlxConn),
	}
}
