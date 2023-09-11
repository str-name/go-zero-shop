package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"

	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryProductListLogic {
	return &CategoryProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryProductListLogic) CategoryProductList(in *pb.CategoryProductListReq) (*pb.CategoryProductListResp, error) {
	// todo: add your logic here and delete this line

	// 判断分类是否存在
	var category model.Category
	err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.CategoryID, globalKey.DelStateNo).Take(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_CATEGORY_NOT_EXISTS_ERROR), " categoryID: %v", in.CategoryID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`category ERROR: %+v", err)
	}

	offset, limit := int((in.Page-1)*in.Size), int(in.Size)
	var list []model.Product
	err = l.svcCtx.ProductDB.Where("category_id = ? and del_state = ? and on_sale = ?", in.CategoryID, globalKey.DelStateNo, globalKey.ProductOnline).
		Order(in.Sort).Offset(offset).Limit(limit).Find(&list).Error
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

	return &pb.CategoryProductListResp{SmallProducts: res}, nil
}
