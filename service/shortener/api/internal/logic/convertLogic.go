// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"errors"
	"mall/service/shortener/pkg/connect"
	"mall/service/shortener/pkg/md5"
	"mall/service/shortener/pkg/urltool"

	"mall/service/shortener/api/internal/svc"
	"mall/service/shortener/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// todo: add your logic here and delete this line
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效链接")
	}
	md5value := md5.SumMd5([]byte(req.LongUrl))
	_, error := l.svcCtx.ShortUrlMapModel.FindOneByMd5(l.ctx, sql.NullString{String: md5value, Valid: true})
	if error != sqlx.ErrNotFound {
		if error == nil {
			return nil, errors.New("该链接已被转过")
		}
		return nil, error
	}
	basePath, err := urltool.GetBasePath(req.LongUrl)

	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlMapModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("该链接已经是短链了")
		}
	}
	return
}
