package userInfo

import (
	"context"
	"zero-shop/common/ctxData"
	"zero-shop/common/tool"
	"zero-shop/user/rpc/user"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMoneyLogic {
	return &GetUserMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMoneyLogic) GetUserMoney(req *types.GetUserMoneyReq) (resp *types.GetUserMoneyResp, err error) {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	moneyResp, err := l.svcCtx.UserRpc.GetUserMoney(l.ctx, &user.GetUserMoneyReq{
		UserID:   userID,
		Password: req.Password,
	})

	resp = new(types.GetUserMoneyResp)
	resp.Money = tool.FenToYuan(moneyResp.Money)

	return
}
