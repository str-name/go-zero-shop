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

type ShelfProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShelfProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShelfProductLogic {
	return &ShelfProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShelfProductLogic) ShelfProduct(in *pb.ShelfProductReq) (*pb.ShelfProductResp, error) {

	// 只有下架的商品才能上架
	var ps []model.Product
	err := l.svcCtx.ProductDB.Where("boss_id = ? and del_state = ? and on_sale = ?", in.BossID, globalKey.DelStateNo, globalKey.ProductNotOnline).
		Where("id in ?", in.ProductIDs).Find(&ps).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR: %+v", err)
	}
	if len(in.ProductIDs) != len(ps) {
		return nil, errors.Wrapf(xerr.NewErrMsg("选择上架的商品有误，请重新选择"), "ShelfProductReq: %+v", in)
	}

	// 批量更新数据
	err = l.svcCtx.ProductDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&ps).Updates(map[string]interface{}{"on_sale": globalKey.ProductOnline}).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATES product ERROR: %+v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ShelfProductResp{}, nil
}
