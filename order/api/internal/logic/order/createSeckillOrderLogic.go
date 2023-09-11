package order

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/ctxData"
	"zero-shop/common/xerr"
	"zero-shop/order/rpc/order"
	"zero-shop/user/rpc/user"

	"zero-shop/order/api/internal/svc"
	"zero-shop/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	return &CreateSeckillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(req *types.CreateSeckillOrderReq) (resp *types.CreateOrderResp, err error) {

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	// 判断用户和地址是否存在
	existsResp, err := l.svcCtx.UserRpc.CheckUserAndAddressExists(l.ctx, &user.CheckUserAndAddressExistsReq{
		UserID:    userID,
		AddressID: req.UserAddressID,
	})
	if err != nil {
		return nil, err
	} else if !existsResp.IsExists {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_AND_ADDRESS_NOT_EXISTS_ERROR),
			"USER OR ADDRESS NOT EXISTS, UserID: %v", userID)
	}

	orderResp, err := l.svcCtx.OrderRpc.CreateSeckillOrder(l.ctx, &order.CreateSeckillOrderReq{
		UserID:        userID,
		SeckillID:     req.SeckillID,
		UserAddressID: req.UserAddressID,
		ProductCount:  req.ProductCount,
		Remark:        req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateOrderResp{OrderSn: orderResp.OrderSn}, nil
}
