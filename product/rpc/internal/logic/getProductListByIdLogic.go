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

type GetProductListByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListByIDLogic {
	return &GetProductListByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductListByIDLogic) GetProductListByID(in *pb.GetProductListByIDReq) (*pb.GetProductListByIDResp, error) {
	// todo: add your logic here and delete this line

	var list []model.Product
	err := l.svcCtx.ProductDB.Where("id in ? and del_state = ?", in.IDList, globalKey.DelStateNo).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR: %+v", err)
	}

	var res []*pb.SmallProduct
	for _, p := range list {
		var sm = pb.SmallProduct{
			ID:            p.ID,
			Title:         p.Title,
			Banner:        p.Banner,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
		}
		res = append(res, &sm)
	}

	return &pb.GetProductListByIDResp{ProductList: res}, nil
}
