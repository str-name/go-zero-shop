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

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
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

	// 获取头像Path
	var imgPath string
	err = l.svcCtx.UserDB.Model(&model.UserHeaderImage{}).
		Where("id = ? and del_state = ?", u.HeaderImageID, globalKey.DelStateNo).Select("path").Scan(&imgPath).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE imgPath ERROR: %v", err)
	}

	return &pb.GetUserInfoResp{
		ID:           u.ID,
		Mobile:       u.Mobile,
		Username:     u.Username,
		Email:        u.Email,
		Sex:          u.Sex,
		HeaderImg:    imgPath,
		Signature:    u.Signature,
		Introduction: u.Introduction,
		Money:        u.Money,
	}, nil
}
