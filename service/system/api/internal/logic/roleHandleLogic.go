package logic

import (
	"context"

	"mall/service/system/api/internal/svc"
	"mall/service/system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleHandleLogic {
	return &RoleHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleHandleLogic) RoleHandle(req *types.RoleSearchRequest) (resp *types.RoleSearchResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
