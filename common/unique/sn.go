package unique

import (
	"fmt"
	"time"
	"zero-shop/common/tool"
)

const (
	ORDER_PRODUCT_PREFIX string = "101"
	ORDER_SECKILL_PREFIX string = "201"
	PAYMENT_PREFIX       string = "301"
)

func GenerateSn(prefix string) string {
	return fmt.Sprintf("%s%s%s", prefix, time.Now().Format("20060102150405"), tool.RandNumToString(8))
}
