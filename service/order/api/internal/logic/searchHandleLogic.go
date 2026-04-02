// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"mall/service/order/api/internal/errorx"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHandleLogic {
	return &SearchHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchHandleLogic) SearchHandle(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line

	//u, error := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.UserName)
	//if error == sqlx.ErrNotFound {
	//	logx.Errorw(error.Error())
	//	return &types.LoginResponse{Message: "该账号不存在"}, errors.New("该账号不存在")
	//}
	//if error != nil {
	//	logx.Errorw(error.Error())
	//	return &types.LoginResponse{Message: "该账号不存在"}, errors.New("查询失败")
	//}
	//password := SaltPassword([]byte(req.Password))
	//if password != u.Password {
	//	return &types.LoginResponse{Message: "密码错误"}, errors.New("密码错误")
	//}
	//expire := l.svcCtx.Config.Auth.AccessExpire
	//nowUnix := time.Now().Unix()
	//token, error := l.generateAccessToken(l.svcCtx.Config.Auth.AccessSecret, u.UserId, expire)
	//if error != nil {
	//	logx.Errorw(error.Error())
	//	return nil, errors.New("登陆失败")
	//}
	userResp, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userclient.GetUserReq{UserId: 1775036814})
	if err != nil {
		return nil, errorx.NewCodeError(errorx.QueryError, "查询错误")

	}
	return &types.SearchResponse{OrderId: 1, Status: 200, Username: userResp.GetUserName()}, nil
}
