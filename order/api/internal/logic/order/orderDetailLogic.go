package order

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/order/rpc/order"
	"zero-shop/user/rpc/user"

	"zero-shop/order/api/internal/svc"
	"zero-shop/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.GetOrderDetailReq) (resp *types.GetOrderDetailResp, err error) {
	// todo: add your logic here and delete this line

	// 判断用户是否存在
	existResp, err := l.svcCtx.UserRpc.CheckUserExists(l.ctx, &user.CheckUserExistsReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	if !existResp.IsExists {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "USER NOT EXISTS, UserID: %v", req.UserID)
	}

	orderResp, err := l.svcCtx.OrderRpc.OrderDetail(l.ctx, &order.OrderDetailReq{
		UserID:  req.UserID,
		OrderSn: req.OrderSn,
	})
	if err != nil {
		return nil, err
	}

	var res = &types.GetOrderDetailResp{
		ID:               orderResp.ID,
		CreateTime:       orderResp.CreateTime,
		UpdateTime:       orderResp.UpdateTime,
		OrderSn:          orderResp.OrderSn,
		UserID:           orderResp.UserID,
		AddressDetail:    orderResp.AddressDetail,
		AddressPhoneName: orderResp.AddressPhoneName,
		ProductID:        orderResp.ProductID,
		Title:            orderResp.Title,
		SubTitle:         orderResp.SubTitle,
		Banner:           orderResp.Banner,
		Info:             orderResp.Info,
		ProductStoreID:   orderResp.ProductStoreID,
		ProductBossID:    orderResp.ProductBossID,
		ProductCount:     orderResp.ProductCount,
		UnitPrice:        tool.FenToYuan(orderResp.UnitPrice),
		TotalPrice:       tool.FenToYuan(orderResp.TotalPrice),
		Status:           orderResp.Status,
		Remark:           orderResp.Remark,
		PayTime:          orderResp.PayTime,
		PayType:          orderResp.PayType,
	}

	return res, nil
}
