package userInfo

import (
	"context"
	"zero-shop/common/ctxData"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)

	infoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		UserInfo: types.User{
			ID:           infoResp.ID,
			Mobile:       infoResp.Mobile,
			Username:     infoResp.Username,
			Email:        infoResp.Email,
			Sex:          infoResp.Sex,
			HeaderImg:    infoResp.HeaderImg,
			Signature:    infoResp.Signature,
			Introduction: infoResp.Introduction,
		},
	}, nil
}
