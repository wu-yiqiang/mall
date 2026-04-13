package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/service/system/api/internal/svc"
	"mall/service/system/api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginHandleLogic {
	return &UserLoginHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var secret = []byte("1Sa@34!Fd#65%778sWJFsdQ")

func SaltPassword(password []byte) string {
	h := md5.New()
	h.Write([]byte(password))
	h.Write(secret)
	return hex.EncodeToString(h.Sum(nil))
}
func (l *UserLoginHandleLogic) generateAccessToken(secret string, userId string, expire int64) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"userId": userId,
		"iat":    now.Unix(),
		"exp":    expire,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))
}
func (l *UserLoginHandleLogic) UserLoginHandle(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
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
	expire := l.svcCtx.Config.Auth.AccessExpire
	nowUnix := time.Now().Unix()
	token, error := l.generateAccessToken(l.svcCtx.Config.Auth.AccessSecret, u.UserId, expire)
	if error != nil {
		logx.Errorw(error.Error())
		return nil, errors.New("登陆失败")
	}
	return &types.LoginResponse{Message: "登陆成功", AccessToken: token, AccessExpire: int(nowUnix + expire), RefreshAfter: int(nowUnix + expire/2)}, nil
}
