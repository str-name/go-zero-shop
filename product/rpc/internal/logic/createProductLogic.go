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

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// storeProduct
func (l *CreateProductLogic) CreateProduct(in *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	// todo: add your logic here and delete this line

	// 判断storeID和bossID是否正确
	var s model.ProductStore
	err := l.svcCtx.ProductDB.Where("boss_id = ? and id = ? and del_state = ?", in.BossID, in.StoreID, globalKey.DelStateNo).Take(&s).Error
	if err != nil {
		// 删除已经上传的图片
		go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
			l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_STORE_NOT_EXISTS), "bossID: %v, storeID: %v", in.BossID, in.StoreID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`store ERROR: %+v", err)
	}

	// 判断分类是否存在
	var c model.Category
	err = l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.CategoryID, globalKey.DelStateNo).Take(&c).Error
	if err != nil {
		// 删除已经上传的图片
		go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
			l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_CATEGORY_NOT_EXISTS_ERROR), "categoryID: %v", in.CategoryID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`category ERROR: %+v", err)
	}

	// 不允许标题重复
	var p model.Product
	err = l.svcCtx.ProductDB.Where("title = ? and del_state = ?", in.Title, globalKey.DelStateNo).Take(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		// 删除已经上传的图片
		go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
			l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`category ERROR: %+v", err)
	} else if err == nil {
		// 删除已经上传的图片
		go qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
			l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_EXISTS_ERROR), "Product: %+v", in)
	}

	// 创建商品
	err = l.svcCtx.ProductDB.Create(&model.Product{
		CategoryID:   in.CategoryID,
		Title:        in.Title,
		SubTitle:     in.SubTitle,
		Banner:       in.Banner,
		Introduction: in.Introduction,
		Price:        in.Price,
		OnSale:       in.OnSale,
		Stock:        in.Stock,
		StoreID:      in.StoreID,
		BossID:       in.BossID,
	}).Error
	if err != nil {
		// 删除已经上传的图片
		_ = qiniu.DeleteImageList(in.Banner, l.svcCtx.Config.Qiniu.AccessKey,
			l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE product ERROR: %+v", err)
	}

	return &pb.CreateProductResp{}, nil
}
