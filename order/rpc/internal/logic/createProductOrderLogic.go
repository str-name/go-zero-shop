package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
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

type CreateProductOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductOrderLogic {
	return &CreateProductOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductOrderLogic) CreateProductOrder(in *pb.CreateProductOrderReq) (*pb.CreateOrderResp, error) {
	// 判断商品是否存在
	detail, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &product.ProductDetailReq{ProductID: in.ProductID})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("FAILED TO QUERY product RECORD"),
			"failed to query product record, rpc ProductRpc.ProductDetail fail, productID: %v, ERROR: %+v", in.ProductID, err)
	} else if detail.Product == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "productID: %v", in.ProductID)
	}

	// 创建订单信息
	var order = model.Order{
		OrderSn:        unique.GenerateSn(unique.ORDER_PRODUCT_PREFIX),
		UserID:         in.UserID,
		UserAddressID:  in.UserAddressID,
		ProductID:      detail.Product.ID,
		SeckillID:      0,
		ProductStoreID: detail.Product.StoreID,
		ProductBossID:  detail.Product.BossID,
		ProductCount:   in.ProductCount,
		Status:         globalKey.OrderWaitPay,
		Remark:         in.Remark,
	}

	// 计算商品总价格
	if detail.Product.DiscountPrice != 0 {
		order.UnitPrice = detail.Product.DiscountPrice
		order.TotalPrice = detail.Product.DiscountPrice * in.ProductCount
	} else {
		order.UnitPrice = detail.Product.Price
		order.TotalPrice = detail.Product.DiscountPrice * in.ProductCount
	}

	// 创建订单
	err = l.svcCtx.OrderDB.Create(&order).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE order ERROR: %+v", err)
	}

	// 发送消息给延迟队列
	payload, err := json.Marshal(jobtype.DeferCloseProductOrderPayload{
		UserID:  in.UserID,
		OrderSn: order.OrderSn,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CREATE defer close product`order task json.Marshal Fail, ERROR: %+v", err)
	} else {
		_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseProductOrder, payload),
			asynq.ProcessIn(globalKey.CloseProductOrderTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("CREATE defer close product`order task AsynqClient.Enqueue, OrderSn: %v, ERROR: %+v", order.OrderSn, err)
		}
	}

	return &pb.CreateOrderResp{OrderSn: order.OrderSn}, nil
}
