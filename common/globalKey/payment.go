package globalKey

// 支付类型
var (
	WalletPay int64 = 1
	WxPay     int64 = 2
	AlliPay   int64 = 3
)

// 支付状态
var (
	PayCancel  int64 = -1
	PayWaitPay int64 = 0
	PaySuccess int64 = 1
	PayRefund  int64 = 2
)
