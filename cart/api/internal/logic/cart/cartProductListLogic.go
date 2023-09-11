package cart

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/cart/rpc/cart"
	"zero-shop/common/ctxData"
	"zero-shop/common/xerr"
	"zero-shop/user/rpc/user"

	"zero-shop/cart/api/internal/svc"
	"zero-shop/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartProductListLogic {
	return &CartProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartProductListLogic) CartProductList() (resp *types.CartProductListResp, err error) {
	// todo: add your logic here and delete this line

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

	listResp, err := l.svcCtx.CartRpc.CartProductList(l.ctx, &cart.CartProductListReq{UserID: userID})
	if err != nil {
		return nil, err
	}

	var list []types.CartProduct
	for _, p := range listResp.CartProducts {
		var product = types.CartProduct{
			ID:            p.ID,
			ProductID:     p.ProductID,
			Title:         p.Title,
			Banner:        p.Banner,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			Count:         p.Count,
			Checked:       p.Checked,
		}
		list = append(list, product)
	}

	return &types.CartProductListResp{CartProducts: list}, nil
}
