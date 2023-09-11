package seckillProduct

import (
	"context"
	"zero-shop/common/tool"
	"zero-shop/product/rpc/product"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillListLogic {
	return &SeckillListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillListLogic) SeckillList(req *types.GetSeckillListReq) (resp *types.GetSeckillListResp, err error) {
	// todo: add your logic here and delete this line

	listResp, err := l.svcCtx.ProductRpc.SeckillList(l.ctx, &product.SeckillListReq{
		StartTime: req.StartTime,
		Time:      req.Time,
	})
	if err != nil {
		return nil, err
	}

	var res []types.SmallSeckill
	for _, i := range listResp.SeckillList {
		var ss = types.SmallSeckill{
			SeckillID:    i.SeckillID,
			Title:        i.Title,
			Banner:       i.Banner,
			SeckillPrice: tool.FenToYuan(i.SeckillPrice),
		}
		res = append(res, ss)
	}

	return &types.GetSeckillListResp{SeckillProducts: res}, nil
}
