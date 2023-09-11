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

type CheckProductExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckProductExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckProductExistsLogic {
	return &CheckProductExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// others
func (l *CheckProductExistsLogic) CheckProductExists(in *pb.CheckProductExistsReq) (*pb.CheckProductExistsResp, error) {
	// todo: add your logic here and delete this line

	var product model.Product
	err := l.svcCtx.ProductDB.Where("id = ? and del_state = ?", in.ProductID, globalKey.DelStateNo).Take(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.CheckProductExistsResp{IsExists: false}, nil
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product ERROR: %+v", err)
	}

	return &pb.CheckProductExistsResp{IsExists: true}, nil
}
