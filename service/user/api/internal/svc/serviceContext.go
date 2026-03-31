// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"api/internal/config"
	"mall/service/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
