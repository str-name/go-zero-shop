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

type ProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductDetailLogic) ProductDetail(in *pb.ProductDetailReq) (*pb.ProductDetailResp, error) {
	// todo: add your logic here and delete this line

	var p model.Product
	err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.ProductID, globalKey.DelStateNo).Take(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), " productID: %v", in.ProductID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product ERROR: %+v", err)
	}

	var product = &pb.Product{
		ID:            p.ID,
		CategoryID:    p.CategoryID,
		Title:         p.Title,
		SubTitle:      p.SubTitle,
		Banner:        p.Banner,
		Introduction:  p.Introduction,
		Price:         p.Price,
		DiscountPrice: p.DiscountPrice,
		OnSale:        p.OnSale,
		SellCount:     p.SellCount,
		CommentCount:  p.CommentCount,
		StoreID:       p.StoreID,
		BossID:        p.BossID,
	}

	return &pb.ProductDetailResp{Product: product}, nil
}
