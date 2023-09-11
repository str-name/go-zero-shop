package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	// 全局错误码
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差了，请稍后再试"
	message[REQUEST_PARAM_ERROR] = "请求参数错误"
	message[TOKEN_INVALID_ERROR] = "无效的token"
	message[DB_ERROR] = "数据库错误"
	message[DB_REDIS_ERROR] = "REDIS数据库错误"
	message[DB_UPDATE_ZERO_ERROR] = "数据更新失败"

	// 用户错误码
	message[USER_ROLE_BAN] = "用户被禁止使用该功能"
	message[USER_NOT_BOSS_ERROR] = "用户并非商家，无权操作"
	message[USER_PASSWORD_ERROR] = "用户密码错误，请重新输入"
	message[USER_REPASSWORD_ERROR] = "两次密码输入错误"
	message[USER_EXISTS_ERROR] = "用户已存在"
	message[USER_NOT_EXISTS_ERROR] = "用户不存在，请重新登录"
	message[USER_LOGOUT_ERROR] = "用户已登出，请重新登录"
	message[USER_SEND_EMAIL_ERROR] = "用户发送邮件验证码失败"
	message[USER_EMAIL_CODE_ERROR] = "用户邮箱验证码错误"
	message[USER_ADDR_NOT_EXISTS_ERROR] = "用户收货地址不存在"
	message[USER_AND_ADDRESS_NOT_EXISTS_ERROR] = "用户或者收货地址错误"

	// 商品错误码
	message[PRODUCT_CATEGORY_NOT_EXISTS_ERROR] = "商品分类不存在"
	message[PRODUCT_EXISTS_ERROR] = "商品已存在"
	message[PRODUCT_NOT_EXISTS_ERROR] = "商品不存在"
	message[PRODUCT_FAVORITE_NOT_EXISITS_ERROR] = "收藏的商品不存在"
	message[PRODUCT_STORE_NOT_EXISTS] = "店铺不存在"
	message[PRODUCT_SECKILL_NOT_EXISTS_ERROR] = "秒杀商品不存在"
	message[PRODUCT_SECKILL_EXISTS_ERROR] = "秒杀商品已存在"

	// 订单错误码
	message[ORDER_TYPE_ERROR] = "订单类型错误，请检查错误"
	message[ORDER_NOT_EXISTS_ERROR] = "订单不存在"
	message[ORDER_UPDATE_STATUS_ERROR] = "订单状态冲突"
	message[ORDER_DOUBLE_SECKILL_ERROR] = "秒杀商品禁止重复下单"
	message[ORDER_SECKILL_TIME_ERROR] = "该商品的秒杀活动不在活动期内"
	message[ORDER_SECKILL_STOCK_ERROR] = "秒杀商品库存不足"

	// 购物车错误码
	message[CART_NOT_EXISTS_ERROR] = "购物车信息不存在"
	message[CART_COUNT_OR_CHECK_ERROR] = "购物车信息更改有误"

	// 支付错误码
	message[PAYMENT_MONEY_NOT_ENOUGH_ERROR] = "金额不够，支付失败"
	message[PAYMENT_ORDER_STATUS_ERROR] = "订单状态错误"
	message[PAYMENT_USER_ERROR] = "支付用户信息出错"
	message[PAYMENT_NOT_EXISTS_ERROR] = "支付订单不存在"
}

// MapErrMsg 根据错误代码获取错误信息
func MapErrMsg(errCode uint32) string {
	if v, ok := message[errCode]; ok {
		return v
	} else {
		return message[SERVER_COMMON_ERROR]
	}
}

// IsErrCode 判断错误代码是否为自定义错误代码
func IsErrCode(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
