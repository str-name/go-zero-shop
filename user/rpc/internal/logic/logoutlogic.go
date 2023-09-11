package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"time"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/user/rpc/internal/svc"
	"zero-shop/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *pb.LogoutReq) (*pb.LogoutResp, error) {
	// todo: add your logic here and delete this line

	expire := in.Expire
	now := time.Now().Unix()
	redisKey := fmt.Sprintf("%s%s", globalKey.Logout, in.AccessToken)
	cmd := l.svcCtx.RedisDB.Set(l.ctx, redisKey, in.UserID, time.Duration(expire-now)*time.Second)
	if cmd.Err() != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "LOGOUT REDIS ERROR: %+v", cmd.Err())
	}
	return &pb.LogoutResp{}, nil
}
