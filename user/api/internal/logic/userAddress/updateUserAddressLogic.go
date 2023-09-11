package userAddress

import (
	"context"
	"zero-shop/common/ctxData"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(req *types.UpdateUserAddressReq) error {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)

	_, err := l.svcCtx.UserRpc.UpdateUserAddress(l.ctx, &user.UpdateUserAddressReq{
		UserID: userID,
		Address: &user.Address{
			ID:            req.ID,
			IsDefault:     req.IsDefault,
			Province:      req.Province,
			City:          req.City,
			Region:        req.Region,
			DetailAddress: req.DetailAddress,
			Name:          req.Name,
			Phone:         req.Phone,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
