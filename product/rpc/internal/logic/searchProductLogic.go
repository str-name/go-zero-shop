package logic

import (
	"context"
	"fmt"
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

type SearchProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProductLogic) SearchProduct(in *pb.SearchProductReq) (*pb.SearchProductResp, error) {
	// todo: add your logic here and delete this line

	offset, limit := int((in.Page-1)*in.Size), int(in.Size)
	db := l.svcCtx.ProductDB.Where("del_state = ? and on_sale = ?", globalKey.DelStateNo, in.OnSale).Order(in.Sort).Offset(offset).Limit(limit)
	key := fmt.Sprintf("%%%s%%", in.Keyword)
	db = db.Where("title like ? or sub_title like ? or introduction like ?", key, key, key)

	// 判断是否选择了分类ID
	if in.CategoryID != 0 {
		var category model.Category
		err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.CategoryID, globalKey.DelStateNo).Take(&category).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_CATEGORY_NOT_EXISTS_ERROR), " categoryID: %v", in.CategoryID)
			}
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`category ERROR: %+v", err)
		}
		db.Where("category_id = ?", in.CategoryID)
	}

	var list []model.Product
	err := db.Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product`search ERROR: %+v", err)
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

	return &pb.SearchProductResp{SmallProducts: res}, nil
}
