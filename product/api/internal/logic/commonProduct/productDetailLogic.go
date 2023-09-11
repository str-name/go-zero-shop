package commonProduct

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/tool"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req *types.ProductDetailReq) (resp *types.ProductDetailResp, err error) {
	// todo: add your logic here and delete this line

	page, size := tool.CheckBasePage(req.Page, req.Size)
	proResp, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &product.ProductDetailReq{ProductID: req.ProductID})
	if err != nil {
		return nil, errors.Wrapf(err, "GET product ERROR: %+v", err)
	}
	var pro = types.Product{
		ID:            proResp.Product.ID,
		CategoryID:    proResp.Product.CategoryID,
		Title:         proResp.Product.Title,
		SubTitle:      proResp.Product.SubTitle,
		Banner:        proResp.Product.Banner,
		Introduction:  proResp.Product.Introduction,
		Price:         tool.FenToYuan(proResp.Product.Price),
		DiscountPrice: tool.FenToYuan(proResp.Product.DiscountPrice),
		OnSale:        proResp.Product.OnSale,
		SellCount:     proResp.Product.SellCount,
		CommentCount:  proResp.Product.CommentCount,
		StoreID:       proResp.Product.StoreID,
		BossID:        proResp.Product.BossID,
	}

	comResp, err := l.svcCtx.ProductRpc.ProductCommentList(l.ctx, &product.ProductCommentListReq{
		ProductID: req.ProductID,
		Page:      page,
		Size:      size,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GET product`comments ERROR: %+v", err)
	}
	var clist []types.Comment
	for _, c := range comResp.Comments {
		var comment = types.Comment{
			ID:        c.ID,
			UserID:    c.UserID,
			ProductID: c.ProductID,
			IsGood:    c.IsGood,
			Content:   c.Content,
		}
		clist = append(clist, comment)
	}

	return &types.ProductDetailResp{
		Product:  pro,
		Comments: clist,
	}, nil
}
