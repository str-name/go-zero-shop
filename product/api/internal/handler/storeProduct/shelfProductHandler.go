package storeProduct

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/product/api/internal/logic/storeProduct"
	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"
)

func ShelfProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShelfProductReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := storeProduct.NewShelfProductLogic(r.Context(), svcCtx)
		err := l.ShelfProduct(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
