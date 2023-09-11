package globalKey

import "time"

// 订单类型
var (
	OrderProduct = "order_product"
	OrderSeckill = "order_seckill"
)

// 订单保持时间
var (
	CloseProductOrderTimeMinutes time.Duration = 30
	CloseSeckillOrderTimeMinutes time.Duration = 3
)

// 订单状态
var (
	OrderCancel      int64 = -1 // 订单取消
	OrderWaitPay     int64 = 10 // 订单未付款
	OrderPayed       int64 = 20 // 订单已付款
	OrderWaitShip    int64 = 30 // 订单订发货
	OrderWaitReceive int64 = 40 // 订单待收货
	OrderSuccess     int64 = 50 // 订单交易成功
	OrderRefund      int64 = 60 // 订单已退款
)
