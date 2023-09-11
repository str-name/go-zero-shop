package userAddress

import (
	"net/http"

	"zero-shop/common/response"

	"zero-shop/user/api/internal/logic/userAddress"
	"zero-shop/user/api/internal/svc"
)

func UserAddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userAddress.NewUserAddressListLogic(r.Context(), svcCtx)
		resp, err := l.UserAddressList()
		response.HttpResponse(r, w, resp, err)
	}
}
