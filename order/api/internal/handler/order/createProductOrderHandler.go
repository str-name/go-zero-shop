package order

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/order/api/internal/logic/order"
	"zero-shop/order/api/internal/svc"
	"zero-shop/order/api/internal/types"
)

func CreateProductOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateProductOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := order.NewCreateProductOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateProductOrder(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
