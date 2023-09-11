package user

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/xerr"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	if req.Password != req.RePassword {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_REPASSWORD_ERROR), "req: %+v", req)
	}
	rpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Mobile:     req.Mobile,
		Password:   req.Password,
		RePassword: req.RePassword,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.RegisterResp)
	resp.JwtAccess.AccessToken = rpcResp.AccessToken
	resp.JwtAccess.AccessExpire = rpcResp.AccessExpire
	resp.JwtAccess.RefreshAfter = rpcResp.RefreshAfter

	return
}
