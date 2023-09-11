package logic

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"

	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarouselLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarouselLogic {
	return &CarouselLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// commonProduct
func (l *CarouselLogic) Carousel(in *pb.CarouselReq) (*pb.CarouselResp, error) {
	// todo: add your logic here and delete this line

	var list []model.Carousel
	err := l.svcCtx.ProductDB.Where("del_state = ?", globalKey.DelStateNo).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND carousels ERROR: %+v", err)
	}

	var res []*pb.Carousel
	for _, c := range list {
		var carousel = pb.Carousel{
			ProductID: c.ProductID,
			ImgPath:   c.ImgPath,
		}
		res = append(res, &carousel)
	}

	return &pb.CarouselResp{Carousels: res}, nil
}
