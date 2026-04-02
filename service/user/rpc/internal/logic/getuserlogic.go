package logic

import (
	"context"
	"errors"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserRes, error) {
	u, error := l.svcCtx.UserModel.FindOneByUserId(l.ctx, int64(in.UserId))
	if error != nil {
		logx.Errorw(error.Error())
		return &user.GetUserRes{}, errors.New("用户详情查询失败")
	}
	return &user.GetUserRes{UserId: u.UserId, UserName: u.Username}, nil
}
