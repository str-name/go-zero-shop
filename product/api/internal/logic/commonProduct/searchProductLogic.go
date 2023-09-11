package commonProduct

import (
	"context"
	"zero-shop/common/tool"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProductLogic) SearchProduct(req *types.SearchProductReq) (resp *types.SearchProductResp, err error) {
	// todo: add your logic here and delete this line

	page, size, sort := tool.CheckBasePageAndSort(req.Page, req.Size, req.Sort)
	listResp, err := l.svcCtx.ProductRpc.SearchProduct(l.ctx, &product.SearchProductReq{
		Keyword:    req.Keyword,
		Sort:       sort,
		OnSale:     req.OnSale,
		CategoryID: req.CategoryID,
		Page:       page,
		Size:       size,
	})
	if err != nil {
		return nil, err
	}

	var res []types.SmallProduct
	for _, p := range listResp.SmallProducts {
		var sp = types.SmallProduct{
			ID:            p.ID,
			Title:         p.Title,
			Banner:        p.Banner,
			Price:         tool.FenToYuan(p.Price),
			DiscountPrice: tool.FenToYuan(p.DiscountPrice),
		}
		res = append(res, sp)
	}

	return &types.SearchProductResp{Products: res}, nil
}
