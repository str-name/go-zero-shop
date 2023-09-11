package logic

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/cart/db/model"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/product/rpc/product"

	"zero-shop/cart/rpc/internal/svc"
	"zero-shop/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToCartLogic {
	return &AddProductToCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductToCartLogic) AddProductToCart(in *pb.AddProductToCartReq) (*pb.AddProductToCartResp, error) {
	// todo: add your logic here and delete this line

	// 判断商品是否存在
	productResp, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &product.ProductDetailReq{ProductID: in.ProductID})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("CartRpc USE ProductRpc Error"), "CartRpc USE ProductRpc ERROR: %+v", err)
	}
	if err == nil && productResp.Product == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "ProductID: %v", in.ProductID)
	}

	// 创建购物车
	err = l.svcCtx.CartDB.Create(&model.Cart{
		UserID:    in.UserID,
		ProductID: in.ProductID,
		Count:     in.Count,
		Checked:   globalKey.CartNotCheck,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE cart ERROR: %+v", err)
	}

	return &pb.AddProductToCartResp{}, nil
}
