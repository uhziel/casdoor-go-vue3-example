package logic

import (
	"context"
	"errors"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	errNameTypeInvalid = errors.New("name type invalid")
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail() (resp *types.UserDetailResp, err error) {
	name, ok := l.ctx.Value("name").(string)
	if !ok {
		l.Errorw(errNameTypeInvalid.Error())
		err = errNameTypeInvalid
		return
	}

	resp = &types.UserDetailResp{
		Name: name,
	}

	return
}
