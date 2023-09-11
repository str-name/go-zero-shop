package payment

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/ctxData"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/payment/rpc/payment"
	"zero-shop/user/rpc/user"

	"zero-shop/payment/api/internal/svc"
	"zero-shop/payment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayLogic {
	return &OrderPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderPayLogic) OrderPay(req *types.OrderPayReq) (resp *types.OrderPayResp, err error) {

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	// 判断用户是否存在
	existResp, err := l.svcCtx.UserRpc.CheckUserExists(l.ctx, &user.CheckUserExistsReq{UserID: userID})
	if err != nil {
		return nil, err
	}
	if !existResp.IsExists {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "USER NOT EXISTS, UserID: %v", userID)
	}

	payResp, err := l.svcCtx.PaymentRpc.OrderPayment(l.ctx, &payment.OrderPaymentReq{
		UserID:      userID,
		OrderSn:     req.OrderSn,
		ServiceType: req.ServiceType,
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderPayResp{
		PayTotalPrice: tool.FenToYuan(payResp.PayTotalPrice),
		PaySn:         payResp.PaySn,
	}, nil
}
