package jobtype

type DeferCloseProductOrderPayload struct {
	UserID  int64
	OrderSn string
}

type DeferCloseSeckillOrderPayload struct {
	UserID       int64
	OrderSn      string
	RedisKey     string
	RedisField   string
	SeckillCount int64
}
