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

type SoldoutProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSoldoutProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SoldoutProductLogic {
	return &SoldoutProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SoldoutProductLogic) SoldoutProduct(in *pb.SoldoutProductReq) (*pb.SoldoutProductResp, error) {
	// todo: add your logic here and delete this line

	// 只有上架的商品才能上架
	var ps []model.Product
	err := l.svcCtx.ProductDB.Where("boss_id = ? and del_state = ? and on_sale = ?", in.BossID, globalKey.DelStateNo, globalKey.ProductOnline).
		Where("id in ?", in.ProductIDs).Find(&ps).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR: %+v", err)
	}
	if len(in.ProductIDs) != len(ps) {
		return nil, errors.Wrapf(xerr.NewErrMsg("选择下架的商品有误，请重新选择"), "SoldoutProductReq: %+v", in)
	}

	// 批量更新数据
	err = l.svcCtx.ProductDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&ps).Updates(map[string]interface{}{"on_sale": globalKey.ProductNotOnline}).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATES product ERROR: %+v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.SoldoutProductResp{}, nil
}
