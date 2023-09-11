package userProduct

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/ctxData"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectProductListLogic {
	return &CollectProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectProductListLogic) CollectProductList() (resp *types.CollectProductListResp, err error) {
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

	listResp, err := l.svcCtx.ProductRpc.CollectProductList(l.ctx, &product.CollectProductListReq{UserID: userID})
	if err != nil {
		return nil, err
	}

	var res []types.SmallProduct
	for _, p := range listResp.Products {
		var sp = types.SmallProduct{
			ID:            p.ID,
			Title:         p.Title,
			Banner:        p.Banner,
			Price:         tool.FenToYuan(p.Price),
			DiscountPrice: tool.FenToYuan(p.DiscountPrice),
		}
		res = append(res, sp)
	}

	return &types.CollectProductListResp{Products: res}, nil
}
