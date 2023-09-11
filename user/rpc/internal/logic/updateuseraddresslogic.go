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

type UpdateUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(in *pb.UpdateUserAddressReq) (*pb.UpdateUserAddressResp, error) {
	// todo: add your logic here and delete this line

	// 判断用户是否存在
	var u model.User
	err := l.svcCtx.UserDB.Where("id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "userID: %v", in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "gorm take user ERROR: %v", err)
	}

	// 判断地址是否存在
	var addr model.UserAddress
	err = l.svcCtx.UserDB.Where("id = ? and del_state = ?", in.Address.ID, globalKey.DelStateNo).Take(&addr).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_ADDR_NOT_EXISTS_ERROR), "address: %+v", in.Address)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user`Address ERROR: %v, userID: %v", err, in.UserID)
	}

	// 如果创建时，选择了默认地址则需要判断是否已经有默认地址存在
	if in.Address.IsDefault == 1 {
		var addr model.UserAddress
		err = l.svcCtx.UserDB.Where("user_id = ? and del_state = ? and is_default = ?", in.UserID, globalKey.DelStateNo, 1).Take(&addr).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user`adress ERROR: %+v", err)
		}
		if err == nil && addr.ID != in.Address.ID {
			return nil, xerr.NewErrMsg("已经有默认地址，请重新设置")
		}
	}

	// IsDefault会有0值更新的情况，所以不能使用model.UserAddress进行更新，这样会导致0值无法更新，所以要使用map更新
	data := l.generateUpdateMap(in)
	// 更新地址信息
	err = l.svcCtx.UserDB.Model(&addr).Updates(data).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_ZERO_ERROR), "MYSQL UPDATE user`address ERROR: %+v, address: %+v", err, in.Address)
	}

	return &pb.UpdateUserAddressResp{}, nil
}

func (l *UpdateUserAddressLogic) generateUpdateMap(in *pb.UpdateUserAddressReq) map[string]interface{} {
	data := make(map[string]interface{})

	data["is_default"] = in.Address.IsDefault
	if in.Address.Province != "" {
		data["province"] = in.Address.Province
	}
	if in.Address.City != "" {
		data["city"] = in.Address.City
	}
	if in.Address.Region != "" {
		data["region"] = in.Address.Region
	}
	if in.Address.DetailAddress != "" {
		data["detail_address"] = in.Address.DetailAddress
	}
	if in.Address.Name != "" {
		data["name"] = in.Address.Name
	}
	if in.Address.Phone != "" {
		data["phone"] = in.Address.Phone
	}

	return data
}
