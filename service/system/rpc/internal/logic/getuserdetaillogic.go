package logic

import (
	"context"

	"mall/service/system/rpc/internal/svc"
	"mall/service/system/rpc/types/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDetailLogic) GetUserDetail(in *system.GetUserDetailsRequest) (*system.GetUserDetailsResponse, error) {
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &system.GetUserDetailsResponse{UserId: u.UserId, UserName: u.Username, Email: u.Email.String, Gender: u.Gender}, nil
}
