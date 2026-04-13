package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRelationRoleModel = (*customUserRelationRoleModel)(nil)

type (
	// UserRelationRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRelationRoleModel.
	UserRelationRoleModel interface {
		userRelationRoleModel
	}

	customUserRelationRoleModel struct {
		*defaultUserRelationRoleModel
	}
)

// NewUserRelationRoleModel returns a model for the database table.
func NewUserRelationRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserRelationRoleModel {
	return &customUserRelationRoleModel{
		defaultUserRelationRoleModel: newUserRelationRoleModel(conn, c, opts...),
	}
}
