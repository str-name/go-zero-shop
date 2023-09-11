package commonProduct

import (
	"context"
	"zero-shop/common/tool"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendLogic) Recommend() (resp *types.RecommendProductResp, err error) {
	// todo: add your logic here and delete this line

	listResp, err := l.svcCtx.ProductRpc.Recommend(l.ctx, &product.RecommendReq{})
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

	return &types.RecommendProductResp{Products: res}, nil
}
