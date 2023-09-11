package userInfo

import (
	"context"
	"zero-shop/common/ctxData"
	"zero-shop/common/xerr"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq) error {
	// todo: add your logic here and delete this line

	// 判断密码是否相等
	if req.OldPassword == req.NewPassword {
		return xerr.NewErrMsg("新密码不能和旧密码相同")
	}
	if req.NewPassword != req.RePassword {
		return xerr.NewErrCode(xerr.USER_REPASSWORD_ERROR)
	}

	// 获取用户id
	userID := ctxData.GetUserIDFromCtx(l.ctx)

	_, err := l.svcCtx.UserRpc.UpdatePassword(l.ctx, &user.UpdatePasswordReq{
		UserID:      userID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
		RePassword:  req.RePassword,
	})
	if err != nil {
		return err
	}

	return nil
}
