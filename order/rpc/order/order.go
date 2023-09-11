// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package order

import (
	"context"

	"zero-shop/order/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateOrderResp        = pb.CreateOrderResp
	CreateProductOrderReq  = pb.CreateProductOrderReq
	CreateSeckillOrderReq  = pb.CreateSeckillOrderReq
	DeleteOrderReq         = pb.DeleteOrderReq
	DeleteOrderResp        = pb.DeleteOrderResp
	GetOrderOnlyDetailReq  = pb.GetOrderOnlyDetailReq
	GetOrderOnlyDetailResp = pb.GetOrderOnlyDetailResp
	OrderDetailReq         = pb.OrderDetailReq
	OrderDetailResp        = pb.OrderDetailResp
	OrderListReq           = pb.OrderListReq
	OrderListResp          = pb.OrderListResp
	SmallOrder             = pb.SmallOrder
	UpdateOrderStatusReq   = pb.UpdateOrderStatusReq
	UpdateOrderStatusResp  = pb.UpdateOrderStatusResp

	Order interface {
		CreateProductOrder(ctx context.Context, in *CreateProductOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
		CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error)
		OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error)
		OrderDetail(ctx context.Context, in *OrderDetailReq, opts ...grpc.CallOption) (*OrderDetailResp, error)
		DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderResp, error)
		UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error)
		// other
		GetOrderOnlyDetail(ctx context.Context, in *GetOrderOnlyDetailReq, opts ...grpc.CallOption) (*GetOrderOnlyDetailResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) CreateProductOrder(ctx context.Context, in *CreateProductOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.CreateProductOrder(ctx, in, opts...)
}

func (m *defaultOrder) CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderReq, opts ...grpc.CallOption) (*CreateOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.CreateSeckillOrder(ctx, in, opts...)
}

func (m *defaultOrder) OrderList(ctx context.Context, in *OrderListReq, opts ...grpc.CallOption) (*OrderListResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.OrderList(ctx, in, opts...)
}

func (m *defaultOrder) OrderDetail(ctx context.Context, in *OrderDetailReq, opts ...grpc.CallOption) (*OrderDetailResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.OrderDetail(ctx, in, opts...)
}

func (m *defaultOrder) DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.DeleteOrder(ctx, in, opts...)
}

func (m *defaultOrder) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.UpdateOrderStatus(ctx, in, opts...)
}

// other
func (m *defaultOrder) GetOrderOnlyDetail(ctx context.Context, in *GetOrderOnlyDetailReq, opts ...grpc.CallOption) (*GetOrderOnlyDetailResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.GetOrderOnlyDetail(ctx, in, opts...)
}
