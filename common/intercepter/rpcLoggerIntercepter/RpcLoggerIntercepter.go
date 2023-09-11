package rpcLoggerIntercepter

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-shop/common/xerr"
)

func RpcLoggerIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err) // 获取err类型
		if e, ok := causeErr.(*xerr.CodeError); ok {
			// 为自定义错误类型 需要转换成grpc错误
			logx.WithContext(ctx).Errorf("[ RPC ERR ] : %+v", err)
			// 转成grpc 错误
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		} else {
			// grpc 错误
			logx.WithContext(ctx).Errorf("[ RPC ERR ] : %+v", err)
		}
	}
	return resp, err
}
