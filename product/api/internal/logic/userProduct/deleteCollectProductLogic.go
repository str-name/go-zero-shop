package userProduct

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/ctxData"
	"zero-shop/common/xerr"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCollectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectProductLogic {
	return &DeleteCollectProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCollectProductLogic) DeleteCollectProduct(req *types.DeleteCollectProductReq) error {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)
	// 判断用户是否存在
	existResp, err := l.svcCtx.UserRpc.CheckUserExists(l.ctx, &user.CheckUserExistsReq{UserID: userID})
	if err != nil {
		return err
	}
	if !existResp.IsExists {
		return errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "USER NOT EXISTS, UserID: %v", userID)
	}

	_, err = l.svcCtx.ProductRpc.DeleteCollectProduct(l.ctx, &product.DeleteCollectProductReq{
		UserID:    userID,
		ProductID: req.ProductID,
	})
	if err != nil {
		return err
	}

	return nil
}
