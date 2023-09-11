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

type CheckSeckillExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckSeckillExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSeckillExistsLogic {
	return &CheckSeckillExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckSeckillExistsLogic) CheckSeckillExists(in *pb.CheckSeckillExistsReq) (*pb.CheckSeckillExistsResp, error) {
	// todo: add your logic here and delete this line

	var seckill model.SeckillProduct
	err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.SeckillID, globalKey.DelStateNo).Take(&seckill).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.CheckSeckillExistsResp{IsExists: false}, nil
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product`seckill ERROR: %+v", err)
	}

	return &pb.CheckSeckillExistsResp{IsExists: true}, nil
}
