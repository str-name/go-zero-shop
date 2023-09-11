package user

import (
	"context"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.LoginResp)
	resp.JwtAccess.AccessToken = loginResp.AccessToken
	resp.JwtAccess.AccessExpire = loginResp.AccessExpire
	resp.JwtAccess.RefreshAfter = loginResp.RefreshAfter

	return
}
