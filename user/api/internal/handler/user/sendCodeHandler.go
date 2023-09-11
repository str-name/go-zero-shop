package user

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/user/api/internal/logic/user"
	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"
)

func SendCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := user.NewSendCodeLogic(r.Context(), svcCtx)
		err := l.SendCode(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
