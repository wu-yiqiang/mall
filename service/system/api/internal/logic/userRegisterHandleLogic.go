package logic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/service/system/api/internal/svc"
	"mall/service/system/api/internal/types"
	"mall/service/system/model"
)

type UserRegisterHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterHandleLogic {
	return &UserRegisterHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterHandleLogic) UserRegisterHandle(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	u, error := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if error != nil && error != sqlx.ErrNotFound {
		logx.Errorw(error.Error())
		return nil, errors.New("用户查询失败")
	}
	if u != nil {
		return nil, errors.New("该用户已存在")
	}
	saltpassword := SaltPassword([]byte(req.Password))
	userId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	user := &model.User{
		UserId:   userId.String(),
		Username: req.UserName,
		Password: saltpassword,
	}
	_, err = l.svcCtx.UserModel.Insert(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{Message: "注册成功"}, nil
}
