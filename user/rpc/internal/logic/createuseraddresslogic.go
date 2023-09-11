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

type CreateUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserAddressLogic {
	return &CreateUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserAddressLogic) CreateUserAddress(in *pb.CreateUserAddressReq) (*pb.CreateUserAddressResp, error) {
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
	// 如果创建时，选择了默认地址则需要判断是否已经有默认地址存在
	if in.Address.IsDefault == 1 {
		var addr model.UserAddress
		err = l.svcCtx.UserDB.Where("user_id = ? and del_state = ? and is_default = ?", in.UserID, globalKey.DelStateNo, 1).Take(&addr).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user`adress ERROR: %+v", err)
		}
		if err == nil {
			return nil, xerr.NewErrMsg("已经有默认地址，请重新设置")
		}
	}
	// 创建地址
	err = l.svcCtx.UserDB.Create(&model.UserAddress{
		UserID:        in.UserID,
		IsDefault:     in.Address.IsDefault,
		Province:      in.Address.Province,
		City:          in.Address.City,
		Region:        in.Address.Region,
		DetailAddress: in.Address.DetailAddress,
		Name:          in.Address.Name,
		Phone:         in.Address.Phone,
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
			"MYSQL CREATE address ERROR: %+v, userID: %v, address: %+v", err, in.UserID, in.Address)
	}

	return &pb.CreateUserAddressResp{}, nil
}
