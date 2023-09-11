package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/user/db/model"

	"zero-shop/user/rpc/internal/svc"
	"zero-shop/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogic {
	return &UpdateEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmailLogic) UpdateEmail(in *pb.UpdateEmailReq) (*pb.UpdateEmailResp, error) {
	// todo: add your logic here and delete this line

	// 判断用户是否存在
	var u model.User
	err := l.svcCtx.UserDB.Where("id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "userID: %v", in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR: %+v", err)
	}

	// 判断密码是否正确
	if tool.Md5ToString(in.Password) != u.Password {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_PASSWORD_ERROR), "password: %v", in.Password)
	}

	// 判断邮箱验证码是否正确
	redisKey := fmt.Sprintf("%s%d_%s", globalKey.SendCode, in.UserID, in.Email)
	res, err := l.svcCtx.RedisDB.Get(l.ctx, redisKey).Result()
	if err != nil && err.Error() != "redis: nil" {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "Redis ERROR: %+v", err)
	}
	if res != in.Code {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_EMAIL_CODE_ERROR), "code: %v", in.Code)
	}

	// 修改用户邮箱信息
	u.Email = in.Email
	err = l.svcCtx.UserDB.Updates(&u).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_ZERO_ERROR),
			"MYSQL UPDATE user`email ERROR: %+v, userID: %v, email: %v", err, in.UserID, in.Email)
	}

	return &pb.UpdateEmailResp{}, nil
}
