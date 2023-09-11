package ctxData

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	JwtKeyUserID    = "jwtUserID"
	JwtKeyUserState = "jwtUserState"
	JwtKeyIsBoss    = "jwtIsBoss"
	JwtKeyExpire    = "jwtExpire"
)

func GetUserIDFromCtx(ctx context.Context) int64 {
	if jsonUserID, ok := ctx.Value(JwtKeyUserID).(json.Number); ok {
		if userID, err := jsonUserID.Int64(); err == nil {
			return userID
		} else {
			logx.WithContext(ctx).Errorf("GetUserIDFromCtx ERROR: %+v", err)
		}
	}
	return 0
}

func GetUserStateFromCtx(ctx context.Context) int64 {
	if jsonUserState, ok := ctx.Value(JwtKeyUserState).(json.Number); ok {
		if userState, err := jsonUserState.Int64(); err == nil {
			return userState
		} else {
			logx.WithContext(ctx).Errorf("GetUserStateFromCtx ERROR: %+v", err)
		}
	}
	return 0
}

func GetUserIsBossFromCtx(ctx context.Context) int64 {
	if jsonIsBoss, ok := ctx.Value(JwtKeyIsBoss).(json.Number); ok {
		if isBoss, err := jsonIsBoss.Int64(); err == nil {
			return isBoss
		} else {
			logx.WithContext(ctx).Errorf("GetUserIsBossFromCtx ERROR: %+v", err)
		}
	}
	return 0
}

func GetJwtExpireFromCtx(ctx context.Context) int64 {
	if jsonJwtExpire, ok := ctx.Value(JwtKeyExpire).(json.Number); ok {
		if jwtExpire, err := jsonJwtExpire.Int64(); err == nil {
			return jwtExpire
		} else {
			logx.WithContext(ctx).Errorf("GetUserIsBossFromCtx ERROR: %+v", err)
		}
	}
	return 0
}
