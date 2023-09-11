package userInfo

import (
	"net/http"

	"zero-shop/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/user/api/internal/logic/userInfo"
	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"
)

func UploadUserImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadUserImgReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamParseError(r, w, err)
			return
		}

		l := userInfo.NewUploadUserImgLogic(r.Context(), svcCtx, r)
		err := l.UploadUserImg(&req)
		response.HttpResponse(r, w, nil, err)
	}
}
