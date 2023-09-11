package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"

	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillLogic {
	return &CreateSeckillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSeckillLogic) CreateSeckill(in *pb.CreateSeckillReq) (*pb.CreateSeckillResp, error) {
	// todo: add your logic here and delete this line

	t, _ := time.Parse("2006-01-02", in.StartTime)

	// 检查商品是否存在
	var p model.Product
	err := l.svcCtx.ProductDB.Where("id = ? and boss_id = ? and store_id = ? and del_state = ?",
		in.ProductID, in.BossID, in.StoreID, globalKey.DelStateNo).Take(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), " productID: %v", in.ProductID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product ERROR: %+v", err)
	}

	// 判断秒杀商品是否重复
	var s model.SeckillProduct
	err = l.svcCtx.ProductDB.Where("product_id = ? and time = ? and to_days(start_time) = to_days(?)", in.ProductID, in.Time, t).Take(&s).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`seckill ERROR: %+v", err)
	} else if err == nil {
		return nil, xerr.NewErrCode(xerr.PRODUCT_SECKILL_EXISTS_ERROR)
	}

	// 检查库存是否超标
	if in.StockCount > p.Stock {
		return nil, errors.Wrapf(xerr.NewErrMsg("库存数量错误"), "Product Stock ERROR, ProductID: %v", in.ProductID)
	}

	err = l.svcCtx.ProductDB.Create(&model.SeckillProduct{
		ID:           0,
		CreateTime:   time.Time{},
		UpdateTime:   time.Time{},
		DeleleTime:   time.Time{},
		DelState:     0,
		Version:      0,
		ProductID:    in.ProductID,
		StoreID:      in.StoreID,
		SeckillPrice: in.SeckillPrice,
		StockCount:   in.StockCount,
		StartTime:    &t,
		Time:         in.Time,
	}).Error

	return &pb.CreateSeckillResp{}, nil
}
