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

type DeleteOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrderLogic) DeleteOrder(in *pb.DeleteOrderReq) (*pb.DeleteOrderResp, error) {
	// todo: add your logic here and delete this line

	var order model.Order
	err := l.svcCtx.OrderDB.Where("order_sn = ? and user_id = ? and del_state = ?", in.OrderSn, in.UserID, globalKey.DelStateNo).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_EXISTS_ERROR), "userID: %v, orderSn: %v, ERROR: %+v", in.UserID, in.OrderSn, err)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE order ERROR: %+v", err)
	}

	err = l.svcCtx.OrderDB.Model(&order).Updates(map[string]interface{}{
		"del_state": globalKey.DelStateYes,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATE order ERROR: %+v", err)
	}

	return &pb.DeleteOrderResp{}, nil
}
