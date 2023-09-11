package tool

import "zero-shop/common/globalKey"

func CheckBasePage(page, size int64) (int64, int64) {
	p, s := page, size
	if page == 0 {
		p = 1
	}
	if size == 0 {
		s = 10
	}
	return p, s
}

func CheckBasePageAndSort(page, size int64, sort string) (int64, int64, string) {
	p, s, r := page, size, sort
	if page <= 0 {
		p = 1
	}
	if size <= 0 {
		s = 10
	}
	swith := ", update_time desc"
	switch sort {
	case "":
		r = "score desc, update_time desc"
	default:
		r = sort + swith
	}
	return p, s, r
}

func CheckBasePageAndType(page, size int64, t string) (int64, int64, int64) {
	p, s := page, size
	if page <= 0 {
		p = 1
	}
	if size <= 0 {
		s = 10
	}

	var status int64
	switch t {
	case "cancel":
		status = globalKey.OrderCancel
	case "waitPay":
		status = globalKey.OrderWaitPay
	case "payed":
		status = globalKey.OrderPayed
	case "waitShip":
		status = globalKey.OrderWaitShip
	case "waitReceive":
		status = globalKey.OrderWaitReceive
	case "success":
		status = globalKey.OrderSuccess
	case "refund":
		status = globalKey.OrderRefund
	default:
		status = 0
	}
	return p, s, status
}

func CheckCartCountAndCheck(count, check int64) bool {
	if count <= 0 {
		return false
	}
	if check != globalKey.CartChecked && check != globalKey.CartNotCheck {
		return false
	}
	return true
}
