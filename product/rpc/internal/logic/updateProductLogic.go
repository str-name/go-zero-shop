package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/qiniu"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"

	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *pb.UpdateProductReq) (*pb.UpdateProductResp, error) {
	// todo: add your logic here and delete this line

	// 判断更新商品是否存在
	var p model.Product
	err := l.svcCtx.ProductDB.Where("id = ? and category_id = ? and boss_id = ? and del_state = ?",
		in.ProductID, in.CategoryID, in.BossID, globalKey.DelStateNo).Take(&p).Error
	if err != nil {
		if in.Banner != "" {
			// 删除已经上传的图片
			go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
				l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
		}
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "productID: %v", in.ProductID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product ERROR: %+v", err)
	}
	if in.CategoryID != p.CategoryID {
		// 判断分类是否存在
		var c model.Category
		err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.CategoryID, globalKey.DelStateNo).Take(&c).Error
		if err != nil {
			if in.Banner != "" {
				// 删除已经上传的图片
				go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
					l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
			}
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_CATEGORY_NOT_EXISTS_ERROR), "categoryID: %v", in.CategoryID)
			}
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`category ERROR: %+v", err)
		}
	}
	if in.Title != p.Title {
		// 不允许标题重复
		var p1 model.Product
		err := l.svcCtx.ProductDB.Where("title = ? and del_state = ?", in.Title, globalKey.DelStateNo).Take(&p1).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			if in.Banner != "" {
				// 删除已经上传的图片
				go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
					l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
			}
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`category ERROR: %+v", err)
		} else if err == nil {
			if in.Banner != "" {
				// 删除已经上传的图片
				go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
					l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
			}
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_EXISTS_ERROR), "Product: %+v", in)
		}
	}

	// onSale会涉及到0值更新，所以需要使用map进行更新
	updateData := l.getUpdateMap(in, p)
	err = l.svcCtx.ProductDB.Model(&p).Updates(updateData).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATE product ERROR: %+v", err)
	}

	return &pb.UpdateProductResp{}, nil
}

func (l *UpdateProductLogic) getUpdateMap(in *pb.UpdateProductReq, p model.Product) map[string]interface{} {
	var data = make(map[string]interface{})
	if in.CategoryID != p.CategoryID {
		data["category_id"] = in.CategoryID
	}
	if in.Title != p.Title {
		data["title"] = in.Title
	}
	if in.SubTitle != p.SubTitle {
		data["sub_title"] = in.SubTitle
	}
	if in.Banner != "" {
		// 删除原商品上传的图片
		go qiniu.DeleteImageList(p.Banner, l.svcCtx.Config.Qiniu.AccessKey,
			l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
		data["banner"] = in.Banner
	}
	if in.Introduction != p.Introduction {
		data["introduction"] = in.Introduction
	}
	if in.Price != p.Price {
		data["price"] = in.Price
	}
	data["discount_price"] = in.DiscountPrice
	data["on_sale"] = in.OnSale
	data["stock"] = in.Stock

	return data
}
