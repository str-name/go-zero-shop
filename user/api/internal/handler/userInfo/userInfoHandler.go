package userInfo

import (
	"net/http"

	"zero-shop/common/response"
	"zero-shop/user/api/internal/logic/userInfo"
	"zero-shop/user/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userInfo.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.HttpResponse(r, w, resp, err)
	}
}
