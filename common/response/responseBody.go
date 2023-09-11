package response

// ResponseBody 定义返回响应的统一格式
type ResponseBody struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespSuccess(data interface{}) *ResponseBody {
	return &ResponseBody{
		Code: 200,
		Msg:  "OK",
		Data: data,
	}
}

func RespError(errCode uint32, errMsg string) *ResponseBody {
	return &ResponseBody{
		Code: errCode,
		Msg:  errMsg,
	}
}
