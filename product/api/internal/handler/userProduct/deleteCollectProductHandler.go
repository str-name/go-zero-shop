package userProduct

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/product/api/internal/logic/userProduct"
	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"
)

func DeleteCollectProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCollectProductReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := userProduct.NewDeleteCollectProductLogic(r.Context(), svcCtx)
		err := l.DeleteCollectProduct(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
