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

type CheckUserAndAddressExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserAndAddressExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserAndAddressExistsLogic {
	return &CheckUserAndAddressExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserAndAddressExistsLogic) CheckUserAndAddressExists(in *pb.CheckUserAndAddressExistsReq) (*pb.CheckUserAndAddressExistsResp, error) {
	// todo: add your logic here and delete this line

	var u model.User
	err := l.svcCtx.UserDB.Where("id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.CheckUserAndAddressExistsResp{IsExists: false}, nil
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR: %+v", err)
	}
	var address model.UserAddress
	err = l.svcCtx.UserDB.Where("id = ? and user_id = ? and del_state = ?", in.AddressID, in.UserID, globalKey.DelStateNo).Take(&address).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.CheckUserAndAddressExistsResp{IsExists: false}, nil
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user`address ERROR: %+v", err)
	}

	return &pb.CheckUserAndAddressExistsResp{IsExists: true}, nil
}
