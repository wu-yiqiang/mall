// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"errors"
	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailsLogic {
	return &UserDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailsLogic) UserDetails(req *types.UserDetailsRequest) (resp *types.UserDetailsResponse, err error) {
	u, error := l.svcCtx.UserModel.FindOneByUserId(l.ctx, int64(req.UserId))
	if error != nil {
		logx.Errorw(error.Error())
		return &types.UserDetailsResponse{}, errors.New("用户详情查询失败")
	}
	return &types.UserDetailsResponse{UserId: u.UserId, UserName: u.Username, Gender: int(u.Gender), Email: u.Email.String}, nil
}
