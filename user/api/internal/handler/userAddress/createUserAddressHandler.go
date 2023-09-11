package userAddress

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/user/api/internal/logic/userAddress"
	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"
)

func CreateUserAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := userAddress.NewCreateUserAddressLogic(r.Context(), svcCtx)
		err := l.CreateUserAddress(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
