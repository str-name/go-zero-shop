package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/user/db/model"
	"zero-shop/user/rpc/user"

	"zero-shop/user/rpc/internal/svc"
	"zero-shop/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// user
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line

	// 判断手机号是否已经注册
	var u = new(model.User)
	err := l.svcCtx.UserDB.Where("mobile = ? and del_state = ?", in.Mobile, globalKey.DelStateNo).Take(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR: %+v", err)
	} else if u != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_EXISTS_ERROR), "mobile: %v", in.Mobile)
	}

	// 给密码进行加密
	newPassword := tool.Md5ToString(in.Password)

	// 生成默认的用户名
	defaultName := tool.RandAllToString(12)

	// 保存进数据库
	var newUser = model.User{
		Mobile:        in.Mobile,
		Username:      defaultName,
		Password:      newPassword,
		HeaderImageID: globalKey.DefaultUserHeaderImgID,
		Money:         1000000,
		SignStatus:    globalKey.UserStateSites,
		Role:          globalKey.UserRoleComm,
		IsBoss:        globalKey.UserNotBoss,
	}
	err = l.svcCtx.UserDB.Create(&newUser).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE user ERROR: %+v", err)
	}

	// 生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserID: newUser.ID,
		State:  newUser.Role,
		IsBoss: newUser.IsBoss,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "Register Generate TOKEN ERROR")
	}

	return &pb.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
