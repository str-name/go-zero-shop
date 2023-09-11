package commonProduct

import (
	"context"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList() (resp *types.HomePageCategoryResp, err error) {
	// todo: add your logic here and delete this line

	listResp, err := l.svcCtx.ProductRpc.CategoryList(l.ctx, &product.CategoryListReq{})
	if err != nil {
		return nil, err
	}

	var res []types.Category
	for _, c := range listResp.Categories {
		var category = types.Category{
			ID:   c.ID,
			Name: c.Name,
		}
		res = append(res, category)
	}

	return &types.HomePageCategoryResp{CategoryList: res}, nil
}
