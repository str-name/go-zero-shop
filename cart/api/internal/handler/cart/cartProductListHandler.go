package cart

import (
	"net/http"

	"zero-shop/common/response"

	"zero-shop/cart/api/internal/logic/cart"
	"zero-shop/cart/api/internal/svc"
)

func CartProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewCartProductListLogic(r.Context(), svcCtx)
		resp, err := l.CartProductList()
		response.HttpResponse(r, w, resp, err)
	}
}
