package logic

import (
	"context"
	"mall/service/user/model"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {
	// todo: add your logic here and delete this line
	user := &model.User{
		Id:       111,
		Username: req.UserName,
		Password: req.Password,
		Gender:   1,
	}
	l.svcCtx.UserModel.Insert(context.Background(), user)
	return &types.SignupResponse{Message: "success"}, nil
}
