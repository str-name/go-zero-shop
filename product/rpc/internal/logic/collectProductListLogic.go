package logic

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/globalKey"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"
	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectProductListLogic {
	return &CollectProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectProductListLogic) CollectProductList(in *pb.CollectProductListReq) (*pb.CollectProductListResp, error) {
	// todo: add your logic here and delete this line

	var ids []int64
	err := l.svcCtx.ProductDB.Model(model.FavoriteProduct{}).Where("user_id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).
		Select("product_id").Scan(&ids).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL SCAN product`favorite`ids ERROR: %+v", err)
	}

	var list []model.Product
	err = l.svcCtx.ProductDB.Where("id in ? and del_state = ? and on_sale = ?", ids, globalKey.DelStateNo, globalKey.ProductOnline).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR: %+v", err)
	}

	var res []*pb.SmallProduct
	for _, p := range list {
		// 只展示第一张商品图片
		firstBanner := tool.GetFirstImg(p.Banner)
		var sp = pb.SmallProduct{
			ID:            p.ID,
			Title:         p.Title,
			Banner:        firstBanner,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
		}
		res = append(res, &sp)
	}

	return &pb.CollectProductListResp{Products: res}, nil
}
