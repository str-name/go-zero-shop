package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
	"zero-shop/common/ctxData"
	"zero-shop/common/xerr"

	"zero-shop/user/rpc/internal/svc"
	"zero-shop/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	// todo: add your logic here and delete this line

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	token, err := l.getJwtToken(accessSecret, now, accessExpire, in.UserID, in.State, in.IsBoss)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR),
			"getJwtToken ERROR: %+v, userID: %v, State: %v, isBoss: %v", err, in.UserID, in.State, in.IsBoss)
	}

	return &pb.GenerateTokenResp{
		AccessToken:  token,
		AccessExpire: accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userID, state, isBoss int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[ctxData.JwtKeyUserID] = userID
	claims[ctxData.JwtKeyUserState] = state
	claims[ctxData.JwtKeyIsBoss] = isBoss
	claims[ctxData.JwtKeyExpire] = iat + seconds
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
