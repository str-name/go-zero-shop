package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"
	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {
	// todo: add your logic here and delete this line

	var ps []model.Product
	err := l.svcCtx.ProductDB.Where("boss_id = ? and del_state = ?", in.BossID, globalKey.DelStateNo).Where("id in ?", in.ProductIDs).Find(&ps).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR: %+v", err)
	}
	fmt.Println(ps)
	if len(in.ProductIDs) != len(ps) {
		return nil, errors.Wrapf(xerr.NewErrMsg("选择删除的商品有误，请重新选择"), "DeleteProductReq: %+v", in)
	}

	// 批量更新数据
	err = l.svcCtx.ProductDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&ps).Updates(map[string]interface{}{"del_state": globalKey.DelStateYes}).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATES product ERROR: %+v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResp{}, nil
}
