package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
	"zero-shop/common/globalKey"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"
	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecommendLogic) Recommend(in *pb.RecommendReq) (*pb.RecommendResp, error) {
	// todo: add your logic here and delete this line

	// 先从redis中判断是否存储了相关信息
	str, err := l.svcCtx.RedisDB.Get(l.ctx, globalKey.Recommend).Result()
	if err == nil {
		var r []*pb.SmallProduct
		err := json.Unmarshal([]byte(str), &r)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "json.Unmarshal ERROR: %+v", err)
		}
		return &pb.RecommendResp{SmallProducts: r}, nil
	}

	// 获取推荐商品的ID列表
	var ids []int64
	err = l.svcCtx.ProductDB.Model(&model.ProductRecommend{}).Select("product_id").Scan(&ids).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL SCAN product`recommend ERROR: %+v", err)
	}

	var ps []model.Product
	err = l.svcCtx.ProductDB.Where("id in ? and del_state = ? and on_sale = ?", ids, globalKey.DelStateNo, globalKey.ProductOnline).Find(&ps).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR: %+v", err)
	}

	var res []*pb.SmallProduct
	for _, p := range ps {
		// 只展示第一张商品图片
		firstBanner := tool.GetFirstImg(p.Banner)
		var sp = pb.SmallProduct{
			ID:            p.ID,
			Title:         p.Title,
			Banner:        firstBanner,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
		}
		res = append(res, &sp)
	}

	// 将信息保存到redis中
	jsonData, _ := json.Marshal(&res)
	resCmd := l.svcCtx.RedisDB.Set(l.ctx, globalKey.Recommend, string(jsonData), 24*time.Hour)
	if resCmd.Err() != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "REDIS SET product`recommend ERROR: %+v", resCmd.Err())
	}

	return &pb.RecommendResp{SmallProducts: res}, nil
}
