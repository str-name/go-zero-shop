package payment

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/payment/api/internal/logic/payment"
	"zero-shop/payment/api/internal/svc"
	"zero-shop/payment/api/internal/types"
)

func OrderPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderPayReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := payment.NewOrderPayLogic(r.Context(), svcCtx)
		resp, err := l.OrderPay(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
