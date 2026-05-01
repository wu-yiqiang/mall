// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"mall/service/system/api/internal/svc"
	"mall/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailsHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailsHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailsHandleLogic {
	return &UserDetailsHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailsHandleLogic) UserDetailsHandle(req *types.UserDetailsRequest) (resp *types.UserDetailsResponse, err error) {
	u, error := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if error == sqlx.ErrNotFound {
		return &types.UserDetailsResponse{Message: "获取用户详情失败", Data: nil}, nil
	}
	return &types.UserDetailsResponse{Message: "", Data: types.UserDetails{
		UserID:   u.UserId,
		UserName: u.Username,
		Gender:   u.Gender,
	}}, nil
}
