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

type GetUserAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressListLogic {
	return &GetUserAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// userAddress
func (l *GetUserAddressListLogic) GetUserAddressList(in *pb.GetUserAddressListReq) (*pb.GetUserAddressListResp, error) {
	// todo: add your logic here and delete this line

	// 判断用户是否存在
	var u model.User
	err := l.svcCtx.UserDB.Where("id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "userID: %v", in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR: %v", err)
	}

	// 获取列表
	var list []model.UserAddress
	err = l.svcCtx.UserDB.Where("user_id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND address`list ERROR: %+v", err)
	}

	var addresses []*pb.Address
	for _, addr := range list {
		var address = pb.Address{
			ID:            addr.ID,
			IsDefault:     addr.IsDefault,
			Province:      addr.Province,
			City:          addr.City,
			Region:        addr.Region,
			DetailAddress: addr.DetailAddress,
			Name:          addr.Name,
			Phone:         addr.Phone,
		}
		addresses = append(addresses, &address)
	}

	return &pb.GetUserAddressListResp{Addresses: addresses}, nil
}
