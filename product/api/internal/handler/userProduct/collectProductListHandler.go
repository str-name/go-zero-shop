package userProduct

import (
	"net/http"

	"zero-shop/common/response"

	"zero-shop/product/api/internal/logic/userProduct"
	"zero-shop/product/api/internal/svc"
)

func CollectProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userProduct.NewCollectProductListLogic(r.Context(), svcCtx)
		resp, err := l.CollectProductList()
		response.HttpResponse(r, w, resp, err)
	}
}
