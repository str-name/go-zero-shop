package userInfo

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/user/api/internal/logic/userInfo"
	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"
)

func BindEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := userInfo.NewBindEmailLogic(r.Context(), svcCtx)
		err := l.BindEmail(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
