package user

import (
	"context"
	"zero-shop/common/ctxData"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) error {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	expire := ctxData.GetJwtExpireFromCtx(l.ctx)
	_, err := l.svcCtx.UserRpc.Logout(l.ctx, &user.LogoutReq{
		UserID:      userID,
		Expire:      expire,
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return err
	}

	return nil
}
