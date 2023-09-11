// Code generated by goctl. DO NOT EDIT.
package types

type CreateProductOrderReq struct {
	Type          string `json:"type"`
	ProductID     int64  `json:"productId"`
	UserAddressID int64  `json:"userAddressId"`
	ProductCount  int64  `json:"productCount"`
	Remark        string `json:"remark"`
}

type CreateSeckillOrderReq struct {
	SeckillID     int64  `json:"seckillId"`
	UserAddressID int64  `json:"userAddressId"`
	ProductCount  int64  `json:"productCount"`
	Remark        string `json:"remark"`
}

type CreateOrderResp struct {
	OrderSn string `json:"orderSn"`
}

type SmallOrder struct {
	OrderSn    string  `json:"orderSn"`
	Title      string  `json:"title"`
	SubTitle   string  `json:"subTitle"`
	ProductID  int64   `json:"productId"`
	Banner     string  `json:"banner"`
	TotalPrice float64 `json:"totalPrice"`
	Status     string  `json:"status"`
}

type GetOrderListReq struct {
	Page int64  `json:"page"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}

type GetOrderListResp struct {
	OrderList []SmallOrder `json:"orderList"`
}

type GetOrderDetailReq struct {
	UserID  int64  `json:"userId"`
	OrderSn string `json:"orderSn"`
}

type GetOrderDetailResp struct {
	ID               int64   `json:"id"`
	CreateTime       string  `json:"createTime"`
	UpdateTime       string  `json:"updateTime"`
	OrderSn          string  `json:"orderSn"`
	UserID           int64   `json:"userId"`
	AddressDetail    string  `json:"userAddress"`
	AddressPhoneName string  `json:"addressPhoneName"`
	ProductID        int64   `json:"productId"`
	Title            string  `json:"title"`
	SubTitle         string  `json:"subTitle"`
	Banner           string  `json:"banner"`
	Info             string  `json:"info"`
	ProductStoreID   int64   `json:"productStoreId"`
	ProductBossID    int64   `json:"productBossId"`
	ProductCount     int64   `json:"productCount"`
	UnitPrice        float64 `json:"unitPrice"`
	TotalPrice       float64 `json:"totalPrice"`
	Status           int64   `json:"status"`
	Remark           string  `json:"remark"`
	PayTime          string  `json:"payTime"`
	PayType          string  `json:"payType"`
}

type DeleteOrderReq struct {
	OrderSn string `json:"orderSn"`
}