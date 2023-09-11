package user

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"time"
	"zero-shop/common/ctxData"
	"zero-shop/common/globalKey"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) error {
	// todo: add your logic here and delete this line

	// 获取用户id
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	// 获取验证码
	code := tool.RandAllToString(6)

	err := tool.SendEmailCode(l.svcCtx.Config.Email.User, req.Email, code,
		l.svcCtx.Config.Email.Password, l.svcCtx.Config.Email.Host, l.svcCtx.Config.Email.Port)
	if err != nil {
		return err
	}

	// redis中设置验证码
	redisKey := fmt.Sprintf("%s%d_%s", globalKey.SendCode, userID, req.Email)
	err = l.svcCtx.RedisDB.Set(l.ctx, redisKey, code, 300*time.Second).Err()
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS set code ERROR: %+v", err)
	}

	return nil
}
