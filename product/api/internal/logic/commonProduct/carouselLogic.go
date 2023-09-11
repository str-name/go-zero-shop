package commonProduct

import (
	"context"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarouselLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarouselLogic {
	return &CarouselLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CarouselLogic) Carousel() (resp *types.HomePageCarouselResp, err error) {
	// todo: add your logic here and delete this line

	listResp, err := l.svcCtx.ProductRpc.Carousel(l.ctx, &product.CarouselReq{})
	if err != nil {
		return nil, err
	}

	var res []types.Carousel
	for _, c := range listResp.Carousels {
		var carousel = types.Carousel{
			ProductID: c.ProductID,
			ImgPath:   c.ImgPath,
		}
		res = append(res, carousel)
	}

	return &types.HomePageCarouselResp{Carousels: res}, nil
}
