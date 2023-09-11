package userAddress

import (
	"context"
	"zero-shop/common/ctxData"
	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"
	"zero-shop/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddressListLogic {
	return &UserAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddressListLogic) UserAddressList() (resp *types.UserAddressListResp, err error) {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)

	listResp, err := l.svcCtx.UserRpc.GetUserAddressList(l.ctx, &user.GetUserAddressListReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	var res []types.Address
	for _, address := range listResp.Addresses {
		var addr = types.Address{
			ID:            address.ID,
			IsDefault:     address.IsDefault,
			Province:      address.Province,
			City:          address.City,
			Region:        address.Region,
			DetailAddress: address.DetailAddress,
			Name:          address.Name,
			Phone:         address.Phone,
		}
		res = append(res, addr)
	}

	return &types.UserAddressListResp{List: res}, nil
}
