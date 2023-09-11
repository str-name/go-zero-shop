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

type UpdateEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogic {
	return &UpdateEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmailLogic) UpdateEmail(req *types.EmailReq) error {
	// todo: add your logic here and delete this line

	// 判断code是否为空
	if req.Code == "" {
		return xerr.NewErrCode(xerr.USER_EMAIL_CODE_ERROR)
	}

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	_, err := l.svcCtx.UserRpc.UpdateEmail(l.ctx, &user.UpdateEmailReq{
		UserID:   userID,
		Email:    req.Email,
		Password: req.Password,
		Code:     req.Code,
	})
	if err != nil {
		return err
	}

	return nil
}
