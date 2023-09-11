package commonProduct

import (
	"net/http"

	"zero-shop/common/response"

	"zero-shop/product/api/internal/logic/commonProduct"
	"zero-shop/product/api/internal/svc"
)

func CarouselHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := commonProduct.NewCarouselLogic(r.Context(), svcCtx)
		resp, err := l.Carousel()
		response.HttpResponse(r, w, resp, err)
	}
}
