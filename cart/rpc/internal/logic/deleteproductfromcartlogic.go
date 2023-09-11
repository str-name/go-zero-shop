package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/cart/db/model"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"

	"zero-shop/cart/rpc/internal/svc"
	"zero-shop/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductFromCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductFromCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductFromCartLogic {
	return &DeleteProductFromCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductFromCartLogic) DeleteProductFromCart(in *pb.DeleteProductFromCartReq) (*pb.DeleteProductFromCartResp, error) {
	// todo: add your logic here and delete this line

	// 判断购物车是否存在
	var cart model.Cart
	err := l.svcCtx.CartDB.Where("id = ? and user_id = ? and del_state = ?", in.CartID, in.UserID, globalKey.DelStateNo).Take(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.CART_NOT_EXISTS_ERROR), "cartID: %v, userID: %v", in.CartID, in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE cart ERROR: %+v", err)
	}

	// 删除购物车信息
	err = l.svcCtx.CartDB.Model(&cart).Updates(map[string]interface{}{
		"del_state": globalKey.DelStateYes,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATES cart ERROR: %+v", err)
	}

	return &pb.DeleteProductFromCartResp{}, nil
}
