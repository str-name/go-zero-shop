package logic

import (
	"context"
	"github.com/pkg/errors"
	"strings"
	"zero-shop/cart/db/model"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	pb2 "zero-shop/product/rpc/pb"
	"zero-shop/product/rpc/product"

	"zero-shop/cart/rpc/internal/svc"
	"zero-shop/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartProductListLogic {
	return &CartProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartProductListLogic) CartProductList(in *pb.CartProductListReq) (*pb.CartProductListResp, error) {
	// todo: add your logic here and delete this line

	var cartList []model.Cart
	err := l.svcCtx.CartDB.Where("user_id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Find(&cartList).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND cart ERROR: %+v, userID: %v", err, in.UserID)
	}

	// 获取productID列表
	var productIDs []int64
	for _, c := range cartList {
		productIDs = append(productIDs, c.ProductID)
	}

	// 获取products列表
	productList, err := l.svcCtx.ProductRpc.GetProductListByID(l.ctx, &product.GetProductListByIDReq{IDList: productIDs})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("CartRpc USE ProductRpc Fail"),
			"CartRpc USE ProductRpc.GetProductListByID Fail, userID: %v, ERROR: %v", in.UserID, err)
	}

	// 生成product - ID Map
	var productMap = make(map[int64]*pb2.SmallProduct, 0)
	for _, p := range productList.ProductList {
		productMap[p.ID] = p
	}

	// 聚合cart和product
	var list []*pb.CartProduct
	for _, cart := range cartList {
		if product, ok := productMap[cart.ProductID]; ok {
			firstBanner := strings.Split(product.Banner, ",")[0]
			var cp = pb.CartProduct{
				ID:            cart.ID,
				ProductID:     cart.ProductID,
				Title:         product.Title,
				Banner:        firstBanner,
				Price:         product.Price,
				DiscountPrice: product.DiscountPrice,
				Count:         cart.Count,
				Checked:       cart.Checked,
			}
			list = append(list, &cp)
		}
	}

	return &pb.CartProductListResp{CartProducts: list}, nil
}
