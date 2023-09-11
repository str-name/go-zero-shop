package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/user/db/model"

	"zero-shop/user/rpc/internal/svc"
	"zero-shop/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserMoneyLogic {
	return &UpdateUserMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserMoneyLogic) UpdateUserMoney(in *pb.UpdateUserMoneyReq) (*pb.UpdateUserMoneyResp, error) {
	// todo: add your logic here and delete this line

	var user model.User
	err := l.svcCtx.UserDB.Where("id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "userID: %v", in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR: %+v, userID: %v", err, in.UserID)
	}

	err = l.svcCtx.UserDB.Model(&user).Updates(map[string]interface{}{
		"money": in.Money,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATES user`money ERROR: %+v, userID: %v, money: %v", err, in.UserID, in.Money)
	}

	return &pb.UpdateUserMoneyResp{}, nil
}
