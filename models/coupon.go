package models

type Coupon struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Discount float32 `json:"discount"`
	Expire   string  `json:"expire"`
}

type CouponResponse struct {
	Response
	Coupon Coupon `json:"data"`
}
