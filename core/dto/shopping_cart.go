/**
 * Copyright 2013 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-22 21:56
 * description :
 * history :
 */

package dto

type ShoppingCart struct {
	Id         int64  `json:"-"`
	CartKey    string `json:"key"`
	BuyerId    int64  `json:"buyer"`
	Summary    string `json:"summary"`
	UpdateTime int64  `json:"update_time"`
	//Items      []*CartItem `json:"items"`
	TotalNum int     `json:"total_num"` // 总数量
	TotalFee float32 `json:"total"`
	OrderFee float32 `json:"fee"`
	// 运营商
	Vendors []*CartVendorGroup `json:"vendors"`
}

type CartVendorGroup struct {
	VendorId   int64       `json:"vendorId"`
	VendorName string      `json:"vendorName"`
	ShopId     int64       `json:"shopId"`
	ShopName   string      `json:"shopName"`
	Items      []*CartItem `json:"items"`
	//结算商品项数目
	CheckedNum int `json:"checkedNum"`
}

type CartItem struct {
	GoodsId    int64   `json:"skuId"`
	GoodsName  string  `json:"name"`
	GoodsNo    string  `json:"no"`
	SmallTitle string  `json:"title"`
	GoodsImage string  `json:"image"`
	Quantity   int     `json:"num"`
	Price      float32 `json:"price"`
	SalePrice  float32 `json:"salePrice"`
	// 是否结算
	Checked bool `json:"checked"`
}
