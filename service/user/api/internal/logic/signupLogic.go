package logic

import (
	"context"
	"errors"
	"mall/service/user/model"
	"time"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var secret = []byte("1Sa@34!Fd#65%778sWJFsdQ")

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
	u, error := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if error != nil && error != sqlx.ErrNotFound {
		logx.Errorw(error.Error())
		return nil, errors.New("用户查询失败")
	}
	if u != nil {
		return nil, errors.New("该用户已存在")
	}
	saltpassword := SaltPassword([]byte(req.Password))
	user := &model.User{
		UserId:   time.Now().Unix(),
		Username: req.UserName,
		Password: saltpassword,
		Gender:   int64(req.Gender),
	}
	l.svcCtx.UserModel.Insert(context.Background(), user)
	return &types.SignupResponse{Message: "success"}, nil
}
