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

type GetUserAddressDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressDetailLogic {
	return &GetUserAddressDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAddressDetailLogic) GetUserAddressDetail(in *pb.GetUserAddressDetailReq) (*pb.GetUserAddressDetailResp, error) {
	// todo: add your logic here and delete this line

	var address model.UserAddress
	err := l.svcCtx.UserDB.Where("id = ? and del_state = ?", in.AddressID, globalKey.DelStateNo).Take(&address).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_ADDR_NOT_EXISTS_ERROR), "addressID: %v", in.AddressID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user`address ERROR: %+v", err)
	}

	return &pb.GetUserAddressDetailResp{
		ID:            address.ID,
		UserID:        address.UserID,
		IsDefault:     address.IsDefault,
		Province:      address.Province,
		City:          address.City,
		Region:        address.Region,
		DetailAddress: address.DetailAddress,
		Name:          address.Name,
		Phone:         address.Phone,
	}, nil
}
