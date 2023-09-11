package cart

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/cart/api/internal/logic/cart"
	"zero-shop/cart/api/internal/svc"
	"zero-shop/cart/api/internal/types"
)

func DeleteProductFromCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeteleProductFromCartReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := cart.NewDeleteProductFromCartLogic(r.Context(), svcCtx)
		err := l.DeleteProductFromCart(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
