package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/order/db/model"

	"zero-shop/order/rpc/internal/svc"
	"zero-shop/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderOnlyDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderOnlyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderOnlyDetailLogic {
	return &GetOrderOnlyDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// other
func (l *GetOrderOnlyDetailLogic) GetOrderOnlyDetail(in *pb.GetOrderOnlyDetailReq) (*pb.GetOrderOnlyDetailResp, error) {
	// todo: add your logic here and delete this line

	var order model.Order
	err := l.svcCtx.OrderDB.Where("user_id = ? and order_sn = ? and del_state = ?", in.UserID, in.OrderSn, globalKey.DelStateNo).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_EXISTS_ERROR), "userID: %v, orderSn: %v", in.UserID, in.OrderSn)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE order ERROR: %+v, userID: %v, orderSn: %v", err, in.UserID, in.OrderSn)
	}

	return &pb.GetOrderOnlyDetailResp{
		ID:             order.ID,
		CreateTime:     order.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:     order.UpdateTime.Format("2006-01-02 15:04:05"),
		OrderSn:        order.OrderSn,
		UserID:         order.UserID,
		ProductID:      order.ProductID,
		ProductStoreID: order.ProductStoreID,
		ProductBossID:  order.ProductBossID,
		ProductCount:   order.ProductCount,
		UnitPrice:      order.UnitPrice,
		TotalPrice:     order.TotalPrice,
		Status:         order.Status,
		Remark:         order.Remark,
	}, nil
}
