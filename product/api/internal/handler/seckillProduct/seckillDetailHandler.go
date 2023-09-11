package seckillProduct

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/product/api/internal/logic/seckillProduct"
	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"
)

func SeckillDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSeckillDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := seckillProduct.NewSeckillDetailLogic(r.Context(), svcCtx)
		resp, err := l.SeckillDetail(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
