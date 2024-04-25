package models

type OrderFood struct {
	Customer      Customer `json:"customer"`
	ID            int      `json:"id"`
	Menu          string   `json:"menu"`
	Quantity      int      `json:"quantity`
	Price         string   `json:"price"`
	Voucher       string   `json:"voucher"`
	PaymentMethod string   `json:"paymentmethod"`
}

type OrderFoodResponse struct {
	Response
	OrderFood OrderFood `json:"data"`
}
