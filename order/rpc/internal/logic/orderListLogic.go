package logic

import (
	"context"
	"github.com/pkg/errors"
	"strings"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/order/db/model"
	"zero-shop/product/rpc/product"

	"zero-shop/order/rpc/internal/svc"
	"zero-shop/order/rpc/pb"
	pb2 "zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListLogic) OrderList(in *pb.OrderListReq) (*pb.OrderListResp, error) {
	// todo: add your logic here and delete this line

	offset, limit := int((in.Page-1)*in.Size), int(in.Size)
	db := l.svcCtx.OrderDB.Where("user_id = ? and del_state = ?", in.UserID, globalKey.DelStateNo).Offset(offset).Limit(limit)
	if in.Status != 0 {
		db = db.Where("status = ?", in.Status)
	}

	var orderList []model.Order
	err := db.Find(&orderList).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND order ERROR: %+v", err)
	}
	// 获取商品结合ID
	var productIDs []int64
	for _, o := range orderList {
		productIDs = append(productIDs, o.ProductID)
	}
	// 获取商品集合
	plist, err := l.svcCtx.ProductRpc.GetProductListByID(l.ctx, &product.GetProductListByIDReq{IDList: productIDs})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("ORDER USE ProductRpc.GetProductListByID Fail"),
			"ORDER USE ProductRpc.GetProductListByID Fail, IDList: %v, ERROR: %+v", productIDs, err)
	}
	// 创建商品Map，ID和商品的对应关系
	var productMap = make(map[int64]*pb2.SmallProduct)
	for _, p := range plist.ProductList {
		productMap[p.ID] = p
	}

	var res []*pb.SmallOrder
	// 聚合订单和商品信息
	for _, o := range orderList {
		if p, ok := productMap[o.ProductID]; ok {
			// 商品存在
			var so = pb.SmallOrder{
				OrderSn:    o.OrderSn,
				Title:      p.Title,
				SubTitle:   p.Title,
				ProductID:  o.ProductID,
				Banner:     strings.Split(p.Banner, ",")[0],
				TotalPrice: o.TotalPrice,
				Status:     l.getOrderStatus(o.Status),
			}
			res = append(res, &so)
		}
	}
	return &pb.OrderListResp{OrderList: res}, nil
}

func (l *OrderListLogic) getOrderStatus(status int64) string {
	switch status {
	case globalKey.OrderCancel:
		return "cancel"
	case globalKey.OrderRefund:
		return "refund"
	case globalKey.OrderWaitPay:
		return "waitPay"
	case globalKey.OrderPayed:
		return "payed"
	case globalKey.OrderSuccess:
		return "success"
	case globalKey.OrderWaitReceive:
		return "waitReceive"
	case globalKey.OrderWaitShip:
		return "waitShip"
	default:
		return "success"
	}
}
