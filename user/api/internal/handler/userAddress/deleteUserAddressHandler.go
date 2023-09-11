package userAddress

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/user/api/internal/logic/userAddress"
	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"
)

func DeleteUserAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := userAddress.NewDeleteUserAddressLogic(r.Context(), svcCtx)
		err := l.DeleteUserAddress(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
