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

type CreateCollectProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCollectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCollectProductLogic {
	return &CreateCollectProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// userProduct
func (l *CreateCollectProductLogic) CreateCollectProduct(in *pb.CreateCollectProductReq) (*pb.CreateCollectProductResp, error) {
	// todo: add your logic here and delete this line

	// 判断商品是否存在
	var p model.Product
	err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.ProductID, globalKey.DelStateNo).Take(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "productID: %v", in.ProductID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product ERROR: %+v", err)
	}

	// 判断是否已经收藏过商品
	var f model.FavoriteProduct
	err = l.svcCtx.ProductDB.Where("user_id = ? and product_id = ? and del_state = ?",
		in.UserID, in.ProductID, globalKey.DelStateNo).Take(&f).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`favorite ERROR: %+v", err)
	}
	if err == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户已收藏过该商品"), "productID: %v", in.ProductID)
	}

	err = l.svcCtx.ProductDB.Create(&model.FavoriteProduct{
		UserID:       in.UserID,
		ProductID:    in.ProductID,
		ProductTitle: p.Title,
		StoreID:      p.StoreID,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE product`favorite ERROR: %+v", err)
	}

	return &pb.CreateCollectProductResp{}, nil
}
