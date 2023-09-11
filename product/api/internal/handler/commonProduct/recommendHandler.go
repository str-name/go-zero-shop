package commonProduct

import (
	"net/http"

	"zero-shop/common/response"

	"zero-shop/product/api/internal/logic/commonProduct"
	"zero-shop/product/api/internal/svc"
)

func RecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := commonProduct.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend()
		response.HttpResponse(r, w, resp, err)
	}
}
