package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	gorm2 "zero-shop/common/gorm"
	"zero-shop/order/rpc/order"
	"zero-shop/payment/rpc/internal/config"
	"zero-shop/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	PaymentDB *gorm.DB

	UserRpc  user.User
	OrderRpc order.Order

	MqKqOrder *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		PaymentDB: gorm2.PaymentDB,

		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),

		MqKqOrder: kq.NewPusher(c.KqPaymentUpdateOrderStateConf.Brokers, c.KqPaymentUpdateOrderStateConf.Topic),
	}
}
