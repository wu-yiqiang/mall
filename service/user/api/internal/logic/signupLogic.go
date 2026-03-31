// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"fmt"

	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/service/user/model"
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
	fmt.Println("req:%#v\n", req)
	user := &model.User{
		UserId:   666,
		Username: req.UserName,
		Password: req.Password,
		Gender:   1,
	}
	_, err = l.svcCtx.UserModel.Insert(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &types.SignupResponse{Message: "Success"}, nil
}
