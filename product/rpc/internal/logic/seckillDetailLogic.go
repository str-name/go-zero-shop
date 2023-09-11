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

type SeckillDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillDetailLogic {
	return &SeckillDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SeckillDetailLogic) SeckillDetail(in *pb.SeckillDetailReq) (*pb.SeckillDetailResp, error) {

	var seckill model.SeckillProduct
	err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.SeckillID, globalKey.DelStateNo).Take(&seckill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_SECKILL_NOT_EXISTS_ERROR), "seckillID: %v", in.SeckillID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`seckill ERROR: %+v", err)
	}

	var product model.Product
	err = l.svcCtx.ProductDB.Where("id = ? and del_state = ?", seckill.ProductID, globalKey.DelStateNo).Take(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "productID: %v", seckill.ProductID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product ERROR: %+v", err)
	}

	return &pb.SeckillDetailResp{Product: &pb.SeckillProduct{
		ID:            seckill.ID,
		ProductID:     product.ID,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		SubTitle:      product.SubTitle,
		Banner:        product.Banner,
		Introduction:  product.Introduction,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		OnSale:        product.OnSale,
		SellCount:     product.SellCount,
		CommentCount:  product.CommentCount,
		StoreID:       product.StoreID,
		BossID:        product.BossID,
		SeckillPrice:  seckill.SeckillPrice,
		SeckillCount:  seckill.StockCount,
		StartTime:     seckill.StartTime.String(),
		Time:          seckill.Time,
	}}, nil
}
