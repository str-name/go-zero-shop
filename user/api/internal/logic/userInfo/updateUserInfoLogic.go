package userInfo

import (
	"context"
	"zero-shop/common/ctxData"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) error {
	// todo: add your logic here and delete this line

	userID := ctxData.GetUserIDFromCtx(l.ctx)

	_, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &user.UpdateUserInfoReq{
		UserID:       userID,
		Username:     req.Username,
		Signature:    req.Signature,
		Introduction: req.Introduction,
		Sex:          req.Sex,
	})
	if err != nil {
		return err
	}

	return nil
}
