package response

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"zero-shop/common/xerr"
)

// HttpResponse 自定义http返回函数
func HttpResponse(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		responseBody := RespSuccess(resp)
		httpx.WriteJson(w, http.StatusOK, responseBody)
	} else {
		// 错误返回
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := "服务器开小差了，请稍后再试"

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			// FromError 判断是否为grpc错误，如果是ok为true，不是则为false
			if gStatus, ok := status.FromError(err); ok { // 判断是否为grpc错误
				if xerr.IsErrCode(uint32(gStatus.Code())) {
					errCode = uint32(gStatus.Code())
					errMsg = gStatus.Message()
				}
			}
		}
		logx.WithContext(r.Context()).Errorf("[ API ERROR ] : %+v", err)
		httpx.WriteJson(w, http.StatusBadRequest, RespError(errCode, errMsg))
	}
}

// ParamParseError 自定义参数校验错误返回函数
func ParamParseError(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s : %s", xerr.MapErrMsg(xerr.REQUEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, RespError(xerr.REQUEST_PARAM_ERROR, errMsg))
}

// JwtAuthError 自定义Jwt认证失败错误返回函数
func JwtAuthError(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s : %s", xerr.MapErrMsg(xerr.TOKEN_INVALID_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, RespError(xerr.TOKEN_INVALID_ERROR, errMsg))
}
