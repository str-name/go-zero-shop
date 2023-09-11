package seckillProduct

import (
	"context"
	"zero-shop/common/tool"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillDetailLogic {
	return &SeckillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillDetailLogic) SeckillDetail(req *types.GetSeckillDetailReq) (resp *types.GetSeckillDetailResp, err error) {
	// todo: add your logic here and delete this line

	seckillResp, err := l.svcCtx.ProductRpc.SeckillDetail(l.ctx, &product.SeckillDetailReq{SeckillID: req.SeckillID})
	if err != nil {
		return nil, err
	}

	var res = types.SeckillProduct{
		Product: types.Product{
			ID:            seckillResp.Product.ID,
			CategoryID:    seckillResp.Product.CategoryID,
			Title:         seckillResp.Product.Title,
			SubTitle:      seckillResp.Product.SubTitle,
			Banner:        seckillResp.Product.Banner,
			Introduction:  seckillResp.Product.Introduction,
			Price:         tool.FenToYuan(seckillResp.Product.Price),
			DiscountPrice: tool.FenToYuan(seckillResp.Product.DiscountPrice),
			OnSale:        seckillResp.Product.OnSale,
			SellCount:     seckillResp.Product.SellCount,
			CommentCount:  seckillResp.Product.CommentCount,
			StoreID:       seckillResp.Product.StoreID,
			BossID:        seckillResp.Product.BossID,
		},
		SeckillPrice: tool.FenToYuan(seckillResp.Product.SeckillPrice),
		StockCount:   seckillResp.Product.SeckillCount,
		StartTime:    seckillResp.Product.StartTime,
		Time:         seckillResp.Product.Time,
	}

	return &types.GetSeckillDetailResp{SeckillProduct: res}, nil
}
