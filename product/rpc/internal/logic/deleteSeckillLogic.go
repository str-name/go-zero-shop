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

type DeleteSeckillLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillLogic {
	return &DeleteSeckillLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSeckillLogic) DeleteSeckill(in *pb.DeleteSeckillReq) (*pb.DeleteSeckillResp, error) {
	// todo: add your logic here and delete this line

	// 检查秒杀商品是否存在
	var s model.SeckillProduct
	err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.SeckillID, globalKey.DelStateNo).Take(&s).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_SECKILL_NOT_EXISTS_ERROR), "seckillID: %v", in.SeckillID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`seckill ERROR: %+v", err)
	}

	// 更新删除状态
	err = l.svcCtx.ProductDB.Model(&s).Updates(map[string]interface{}{
		"del_state": globalKey.DelStateYes,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATE product`seckill ERROR: %+v", err)
	}

	return &pb.DeleteSeckillResp{}, nil
}
