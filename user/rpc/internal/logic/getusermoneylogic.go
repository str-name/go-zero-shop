package logic

import (
	"context"
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

type GetUserMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMoneyLogic {
	return &GetUserMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserMoneyLogic) GetUserMoney(in *pb.GetUserMoneyReq) (*pb.GetUserMoneyResp, error) {
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

	return &pb.GetUserMoneyResp{Money: u.Money}, nil
}
