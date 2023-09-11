package globalKey

// 用户信息
const (
	Logout   = "logout_"    // 用户登出
	SendCode = "emailCode_" // 邮箱验证码
)

// 商品信息
const (
	Recommend   = "set_product_recommend"
	SeckillList = "set_seckill_list"
)

// 订单信息
const (
	DoubleOrder        = "set_order_seckill_id"     // 秒杀订单重复下单
	DoubleOrderExpire  = 12                         // 秒杀订单重复下单过期时间
	SeckillCount       = "hash_order_seckill_count" // 秒杀商品库存
	SeckillCountExpire = 12                         // 秒杀商品库存过期时间
)
