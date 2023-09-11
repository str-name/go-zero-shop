package commonProduct

import (
	"context"
	"zero-shop/common/tool"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryProductListLogic {
	return &CategoryProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryProductListLogic) CategoryProductList(req *types.CategoryProductListReq) (resp *types.CategoryProductListResp, err error) {
	// todo: add your logic here and delete this line

	page, size, sort := tool.CheckBasePageAndSort(req.Page, req.Size, req.Sort)
	respList, err := l.svcCtx.ProductRpc.CategoryProductList(l.ctx, &product.CategoryProductListReq{
		CategoryID: req.CategoryID,
		Sort:       sort,
		Page:       page,
		Size:       size,
	})
	if err != nil {
		return nil, err
	}

	var res []types.SmallProduct
	for _, p := range respList.SmallProducts {
		var pro = types.SmallProduct{
			ID:            p.ID,
			Title:         p.Title,
			Banner:        p.Banner,
			Price:         tool.FenToYuan(p.Price),
			DiscountPrice: tool.FenToYuan(p.DiscountPrice),
		}
		res = append(res, pro)
	}

	return &types.CategoryProductListResp{Products: res}, nil
}
