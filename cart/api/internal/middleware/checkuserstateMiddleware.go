package middleware

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"net/http"
	"zero-shop/common/ctxData"
	"zero-shop/common/globalKey"
	"zero-shop/common/goredis"
	gorm2 "zero-shop/common/gorm"
	"zero-shop/common/response"
	"zero-shop/common/xerr"
)

type CheckUserStateMiddleware struct {
	mysqlDB *gorm.DB
	redisDB *redis.Client
}

func NewCheckUserStateMiddleware() *CheckUserStateMiddleware {
	return &CheckUserStateMiddleware{
		mysqlDB: gorm2.UserDB,
		redisDB: goredis.Rdb,
	}
}

func (m *CheckUserStateMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// 判断用户是否被禁止
		ctx := r.Context()
		userState := ctxData.GetUserStateFromCtx(ctx)
		if userState == globalKey.UserRoleBan {
			response.HttpResponse(r, w, nil, xerr.NewErrCode(xerr.USER_ROLE_BAN))
			return
		}
		// 检查当前token是否失效
		token := r.Header.Get("Authorization")
		redisExpireKey := fmt.Sprintf("%s%s", globalKey.Logout, token)
		val, err := m.redisDB.Get(ctx, redisExpireKey).Result()
		if err != nil && err.Error() != "redis: nil" {
			logx.WithContext(ctx).Errorf("Middleware REDIS ERROR: %+v", err)
			response.HttpResponse(r, w, nil, xerr.NewErrCode(xerr.DB_REDIS_ERROR))
			return
		}
		if val != "" {
			response.HttpResponse(r, w, nil, xerr.NewErrCode(xerr.USER_LOGOUT_ERROR))
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
