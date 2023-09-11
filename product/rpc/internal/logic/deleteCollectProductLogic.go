package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"

	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCollectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectProductLogic {
	return &DeleteCollectProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCollectProductLogic) DeleteCollectProduct(in *pb.DeleteCollectProductReq) (*pb.DeleteCollectProductResp, error) {
	// todo: add your logic here and delete this line

	// 判断收藏记录是否存在
	var f model.FavoriteProduct
	err := l.svcCtx.ProductDB.Where("user_id = ? and product_id = ? and del_state = ?", in.UserID, in.ProductID, globalKey.DelStateNo).Take(&f).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "userID: %v, productID: %v", in.UserID, in.ProductID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`collect ERROR: %+v", err)
	}

	// 修改del_state
	err = l.svcCtx.ProductDB.Model(&f).Update("del_state", globalKey.DelStateYes).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATE product`collect ERROR: %+v", err)
	}

	return &pb.DeleteCollectProductResp{}, nil
}
