package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/payment/db/model"

	"zero-shop/payment/rpc/internal/svc"
	"zero-shop/payment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentDetailLogic {
	return &GetPaymentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaymentDetailLogic) GetPaymentDetail(in *pb.GetPaymentDetailReq) (*pb.GetPaymentDetailResp, error) {
	// todo: add your logic here and delete this line

	var payment model.Payment
	err := l.svcCtx.PaymentDB.Where("order_sn = ? and del_state = ?", in.OrderSn, globalKey.DelStateNo).Take(&payment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_NOT_EXISTS_ERROR), "orderSn: %v", in.OrderSn)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE payment ERROR: %+v, orderSn: %v", err, in.OrderSn)
	}

	return &pb.GetPaymentDetailResp{
		ID:             payment.ID,
		PaymentSn:      payment.PaymentSn,
		OrderSn:        payment.OrderSn,
		UserID:         payment.UserID,
		PayMode:        payment.PayMode,
		TradeType:      payment.TradeType,
		TradeState:     payment.TradeState,
		PayTotal:       payment.PayTotal,
		TransactionID:  payment.TransactionID,
		TradeStateDesc: payment.TradeStateDesc,
		PayStatus:      payment.PayStatus,
		PayTime:        payment.CreateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
