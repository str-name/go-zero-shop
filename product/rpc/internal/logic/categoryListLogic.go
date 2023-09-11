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

type CategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryListLogic) CategoryList(in *pb.CategoryListReq) (*pb.CategoryListResp, error) {
	// todo: add your logic here and delete this line

	var list []model.Category
	err := l.svcCtx.ProductDB.Where("del_state = ?", globalKey.DelStateNo).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND categories ERROR: %+v", err)
	}

	var res []*pb.Category
	for _, c := range list {
		var category = pb.Category{
			ID:   c.ID,
			Name: c.Name,
		}
		res = append(res, &category)
	}

	return &pb.CategoryListResp{Categories: res}, nil
}
