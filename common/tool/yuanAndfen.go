package tool

import "github.com/shopspring/decimal"

var oneHundredDecimal decimal.Decimal = decimal.NewFromInt(100)

func YuanToFen(yuan float64) int64 {
	f, _ := decimal.NewFromFloat(yuan).Mul(oneHundredDecimal).Truncate(0).Float64()
	return int64(f)
}

func FenToYuan(fen int64) float64 {
	y, _ := decimal.NewFromInt(fen).Div(oneHundredDecimal).Truncate(2).Float64()
	return y
}
