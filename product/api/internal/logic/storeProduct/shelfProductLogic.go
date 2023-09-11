package storeProduct

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

type ShelfProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShelfProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShelfProductLogic {
	return &ShelfProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShelfProductLogic) ShelfProduct(req *types.ShelfProductReq) error {
	// todo: add your logic here and delete this line

	// 获取bossID
	bossID := ctxData.GetUserIDFromCtx(l.ctx)
	// 判断用户是否存在
	existResp, err := l.svcCtx.UserRpc.CheckUserExists(l.ctx, &user.CheckUserExistsReq{UserID: bossID})
	if err != nil {
		return err
	}
	if !existResp.IsExists {
		return errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "USER NOT EXISTS, UserID: %v", bossID)
	}

	_, err = l.svcCtx.ProductRpc.ShelfProduct(l.ctx, &product.ShelfProductReq{
		ProductIDs: req.ProductIDs,
		BossID:     bossID,
	})
	if err != nil {
		return err
	}

	return nil
}
