package cart

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/cart/rpc/cart"
	"zero-shop/common/ctxData"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/user/rpc/user"

	"zero-shop/cart/api/internal/svc"
	"zero-shop/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductDetailLogic {
	return &UpdateProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductDetailLogic) UpdateProductDetail(req *types.UpdateProductDetailReq) error {
	// todo: add your logic here and delete this line

	// 判断数量和选中状态是否正确
	state := tool.CheckCartCountAndCheck(req.Count, req.Check)
	if !state {
		return errors.Wrapf(xerr.NewErrCode(xerr.CART_COUNT_OR_CHECK_ERROR), "count: %v, check: %v", req.Count, req.Check)
	}

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

	_, err = l.svcCtx.CartRpc.UpdateProductDetail(l.ctx, &cart.UpdateProductDetailReq{
		CartID: req.CartID,
		UserID: userID,
		Count:  req.Count,
		Check:  req.Check,
	})
	if err != nil {
		return err
	}

	return nil
}
