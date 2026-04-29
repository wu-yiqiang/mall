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
	var roles []types.Role
	var menus []types.Menu
	queryRoleSql := `SELECT r.role_id, r.name, r.code 
        FROM role r
        JOIN user_relation_role ur ON r.role_id = ur.role_id
        WHERE ur.user_id = ?`
	err = l.svcCtx.DB.QueryRowsCtx(l.ctx, &roles, queryRoleSql, u.UserId)
	if err != nil {
		roles = make([]types.Role, 0)
		return &types.LoginResponse{Message: "登陆失败", AccessToken: "", AccessExpire: 0, RefreshAfter: 0, Roles: roles, Menus: menus}, nil
	}
	queryMenuSql := `select DISTINCT IFNULL(t5.menu_id, "") as parent_id, IFNULL(t5.name, ""), IFNULL(t5.code, ""),IFNULL(t5.path, ""),IFNULL(t5.serial, ""),IFNULL(t5.parent_id, "") as parent_id
	from user_relation_role t1 
join user_relation_role t2 
	on t1.role_id = t2.role_id
join role_relation_menu t4 
on
 t1.role_id = t4.role_id
join menu t5
on t4.menu_id = t5.menu_id
where t1.user_id = ?`
	err = l.svcCtx.DB.QueryRowsCtx(l.ctx, &menus, queryMenuSql, u.UserId)
	if err != nil {
		menus = make([]types.Menu, 0)
		return &types.LoginResponse{Message: "登陆失败", AccessToken: "", AccessExpire: 0, RefreshAfter: 0, Roles: make([]types.Role, 0), Menus: make([]types.Menu, 0)}, nil
	}
	return &types.LoginResponse{Message: "登陆成功", AccessToken: token, AccessExpire: int(nowUnix + expire), RefreshAfter: int(nowUnix + expire/2), Roles: roles, Menus: menus}, nil
}
