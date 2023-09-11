package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"time"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"

	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillListLogic {
	return &SeckillListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// seckillProduct
func (l *SeckillListLogic) SeckillList(in *pb.SeckillListReq) (*pb.SeckillListResp, error) {
	// todo: add your logic here and delete this line

	// 判断redis中是否有数据
	key := fmt.Sprintf("%s_%s_%d", globalKey.SeckillList, in.StartTime, in.Time)
	str, err := l.svcCtx.RedisDB.Get(l.ctx, key).Result()
	if err == nil {
		var sec []*pb.SmallSeckill
		err := json.Unmarshal([]byte(str), &sec)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "json.Unmarshal ERROR: %+v", err)
		}
		return &pb.SeckillListResp{SeckillList: sec}, nil
	}

	t, _ := time.Parse("2006-01-02", in.StartTime)
	// 获取秒杀集合
	var slist []model.SeckillProduct
	err = l.svcCtx.ProductDB.Where("time = ? and to_days(start_time) = to_days(?)", in.Time, t).Find(&slist).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product`seckill ERROR: %+v", err)
	}
	// 获取商品ID列表
	var productIDs []int64
	for _, product := range slist {
		productIDs = append(productIDs, product.ProductID)
	}
	// 获取商品结合
	var plist []model.Product
	err = l.svcCtx.ProductDB.Where("id in ?", productIDs).Find(&plist).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR: %+v", err)
	}
	// 创建商品Map，ID和商品的对应关系
	var productMap = make(map[int64]model.Product)
	for _, p := range plist {
		productMap[p.ID] = p
	}

	seckillStockKey := fmt.Sprintf("%s_%s_%d", globalKey.SeckillCount, in.StartTime, in.Time)
	// 聚合秒杀和商品信息
	var res []*pb.SmallSeckill
	for _, seckillProduct := range slist {
		// 从商品map中获取商品对象
		if p, ok := productMap[seckillProduct.ProductID]; ok {
			// 获取到第一张封面照片
			fisrtBanenr := strings.Split(p.Banner, ",")[0]
			res = append(res, &pb.SmallSeckill{
				SeckillID:    seckillProduct.ID,
				Title:        p.Title,
				Banner:       fisrtBanenr,
				SeckillPrice: seckillProduct.SeckillPrice,
			})
			err := l.svcCtx.RedisDB.HSet(l.ctx, seckillStockKey, seckillProduct.ID, seckillProduct.StockCount).Err()
			if err != nil {
				logx.WithContext(l.ctx).Errorf("REDIS HSET seckillStockKey Fail, Key: %v, ERROR: %+v", seckillStockKey, err)
				continue
			}
			_ = l.svcCtx.RedisDB.Expire(l.ctx, seckillStockKey, globalKey.SeckillCountExpire*time.Hour)
		}
	}

	// 将聚合后的秒杀集合存入redis中
	jsonData, _ := json.Marshal(&res)
	resCmd := l.svcCtx.RedisDB.Set(l.ctx, key, jsonData, 2*time.Hour)
	if resCmd.Err() != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS SET product`seckill ERROR: %+v", resCmd.Err())
	}

	return &pb.SeckillListResp{SeckillList: res}, nil
}
