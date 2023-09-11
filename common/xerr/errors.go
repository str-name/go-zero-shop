package xerr

import "fmt"

type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 返给前端信息
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返给前端信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode: %d, ErrMsg: %s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{
		errCode: errCode,
		errMsg:  errMsg,
	}
}

func NewErrCode(errCode uint32) *CodeError {
	if ok := IsErrCode(errCode); ok {
		return &CodeError{
			errCode: errCode,
			errMsg:  MapErrMsg(errCode),
		}
	} else {
		return &CodeError{
			errCode: SERVER_COMMON_ERROR,
			errMsg:  "服务器开小差了，请稍后再试",
		}
	}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{
		errCode: SERVER_COMMON_ERROR,
		errMsg:  errMsg,
	}
}
