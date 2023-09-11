package cart

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/cart/api/internal/logic/cart"
	"zero-shop/cart/api/internal/svc"
	"zero-shop/cart/api/internal/types"
)

func UpdateProductDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProductDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := cart.NewUpdateProductDetailLogic(r.Context(), svcCtx)
		err := l.UpdateProductDetail(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
