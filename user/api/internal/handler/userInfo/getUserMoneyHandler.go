package userInfo

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/user/api/internal/logic/userInfo"
	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"
)

func GetUserMoneyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserMoneyReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := userInfo.NewGetUserMoneyLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMoney(&req)
		response.HttpResponse(r, w, resp, err)
	}
}
