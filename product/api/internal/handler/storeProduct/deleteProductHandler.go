package storeProduct

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/product/api/internal/logic/storeProduct"
	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"
)

func DeleteProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := storeProduct.NewDeleteProductLogic(r.Context(), svcCtx)
		err := l.DeleteProduct(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
