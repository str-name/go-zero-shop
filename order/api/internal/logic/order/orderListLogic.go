package order

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/ctxData"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/order/rpc/order"
	"zero-shop/user/rpc/user"

	"zero-shop/order/api/internal/svc"
	"zero-shop/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.GetOrderListReq) (resp *types.GetOrderListResp, err error) {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	// 判断用户是否存在
	existResp, err := l.svcCtx.UserRpc.CheckUserExists(l.ctx, &user.CheckUserExistsReq{UserID: userID})
	if err != nil {
		return nil, err
	}
	if !existResp.IsExists {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "USER NOT EXISTS, UserID: %v", userID)
	}

	page, size, status := tool.CheckBasePageAndType(req.Page, req.Size, req.Type)

	listResp, err := l.svcCtx.OrderRpc.OrderList(l.ctx, &order.OrderListReq{
		UserID: userID,
		Page:   page,
		Size:   size,
		Status: status,
	})

	var res []types.SmallOrder
	for _, o := range listResp.OrderList {
		var so = types.SmallOrder{
			OrderSn:    o.OrderSn,
			Title:      o.Title,
			SubTitle:   o.SubTitle,
			ProductID:  o.ProductID,
			Banner:     o.Banner,
			TotalPrice: tool.FenToYuan(o.TotalPrice),
			Status:     o.Status,
		}
		res = append(res, so)
	}

	return &types.GetOrderListResp{OrderList: res}, nil
}
