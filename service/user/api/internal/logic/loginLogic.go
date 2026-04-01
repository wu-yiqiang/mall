// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func SaltPassword(password []byte) string {
	h := md5.New()
	h.Write([]byte(password))
	h.Write(secret)
	return hex.EncodeToString(h.Sum(nil))
}

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	u, error := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	if error == sqlx.ErrNotFound {
		logx.Errorw(error.Error())
		return &types.LoginResponse{Message: "该账号不存在"}, errors.New("该账号不存在")
	}
	if error != nil {
		logx.Errorw(error.Error())
		return &types.LoginResponse{Message: "该账号不存在"}, errors.New("查询失败")
	}
	password := SaltPassword([]byte(req.Password))
	if password != u.Password {
		return &types.LoginResponse{Message: "密码错误"}, errors.New("密码错误")
	}
	return &types.LoginResponse{Message: "登陆成功"}, nil
}
