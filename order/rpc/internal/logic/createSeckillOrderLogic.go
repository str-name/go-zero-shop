package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"time"
	"zero-shop/common/globalKey"
	"zero-shop/common/unique"
	"zero-shop/common/xerr"
	"zero-shop/mqueue/job/jobtype"
	"zero-shop/order/db/model"
	"zero-shop/product/rpc/product"

	"zero-shop/order/rpc/internal/svc"
	"zero-shop/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	return &CreateSeckillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(in *pb.CreateSeckillOrderReq) (*pb.CreateOrderResp, error) {
	// 判断秒杀商品是否存在
	detail, err := l.svcCtx.ProductRpc.SeckillDetail(l.ctx, &product.SeckillDetailReq{SeckillID: in.SeckillID})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("FAILED TO QUERY product`seckill RECORD"),
			"failed to query product`seckill record, rpc ProductRpc.SeckillDetail fail, seckillID: %v, ERROR: %+v", in.SeckillID, err)
	} else if detail.Product == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_SECKILL_NOT_EXISTS_ERROR), "seckillID: %v", in.SeckillID)
	}

	// 判断下单时间是否再秒杀商品活动期间
	startTime := strings.Split(detail.Product.StartTime, " ")[0]
	day := time.Now().Format("2006-01-02")
	hour := int64(time.Now().Hour())
	if day != startTime || hour-detail.Product.Time >= 2 || hour-detail.Product.Time < 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_SECKILL_TIME_ERROR), "seckillID: %v, time: %+v", in.SeckillID, time.Now())
	}

	// 判断用户是否重复下单
	doubleKey := fmt.Sprintf("%s_%d", globalKey.DoubleOrder, in.SeckillID)
	isDouble, err := l.svcCtx.RedisDB.SIsMember(l.ctx, doubleKey, in.UserID).Result()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS SIsMember Fail, key: %v, ERROR: %v", doubleKey, err)
	} else if isDouble {
		// 该用户已经下过单了
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_DOUBLE_SECKILL_ERROR), "userID: %v, seckillID: %v", in.UserID, in.SeckillID)
	}

	// 预查询库存是否充足
	seckillStockKey := fmt.Sprintf("%s_%s_%d", globalKey.SeckillCount, day, detail.Product.Time)
	seckillStockField := fmt.Sprintf("%d", detail.Product.ID)
	stockCmd, err := l.svcCtx.RedisDB.HIncrBy(l.ctx, seckillStockKey, seckillStockField, -in.ProductCount).Result()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR),
			"REDIS HIncrBy Fail, Key: %v, Field: %v, ERROR: %+v", seckillStockKey, seckillStockField, err)
	} else if stockCmd < 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_SECKILL_STOCK_ERROR), "")
	}

	var order = model.Order{
		OrderSn:        unique.GenerateSn(unique.ORDER_SECKILL_PREFIX),
		UserID:         in.UserID,
		UserAddressID:  in.UserAddressID,
		ProductID:      detail.Product.ProductID,
		SeckillID:      in.SeckillID,
		ProductStoreID: detail.Product.StoreID,
		ProductBossID:  detail.Product.BossID,
		ProductCount:   in.ProductCount,
		UnitPrice:      detail.Product.SeckillPrice,
		TotalPrice:     in.ProductCount * detail.Product.SeckillCount,
		Status:         globalKey.OrderWaitPay,
		Remark:         in.Remark,
	}
	err = l.svcCtx.OrderDB.Transaction(func(tx *gorm.DB) error {
		// 创建订单
		err = tx.Create(&order).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE order ERROR: %+v", err)
		}
		// 在Redis中存储userId和seckillId
		err = l.svcCtx.RedisDB.SAdd(l.ctx, doubleKey, in.UserID).Err()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS SADD Fail, key: %v, ERROR: %+v", doubleKey, err)
		}
		_ = l.svcCtx.RedisDB.Expire(l.ctx, doubleKey, globalKey.DoubleOrderExpire*time.Hour)

		// 修改库存
		return nil
	})
	if err != nil {
		return nil, err
	}

	// 发送给延迟队列
	payload, err := json.Marshal(jobtype.DeferCloseSeckillOrderPayload{
		UserID:  in.UserID,
		OrderSn: order.OrderSn,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CREATE defer close seckill`order task json.Marshal Fail, ERROR: %+v", err)
	} else {
		_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseSeckillOrder, payload),
			asynq.ProcessIn(globalKey.CloseSeckillOrderTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("CREATE defer close seckill`order task AsynqClient.Enqueue, OrderSn: %v, ERROR: %+v", order.OrderSn, err)
		}
	}

	return &pb.CreateOrderResp{OrderSn: order.OrderSn}, nil
}
