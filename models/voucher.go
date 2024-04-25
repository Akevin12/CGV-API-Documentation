package models

type Voucher struct {
	Customer    Customer `json:"customer"`
	VoucherName string   `json:"vouchername"`
	Description string   `json:"description"`
	ExpiredDate string   `json:"expireddate`
}

type VoucherResponse struct {
	Response
	Voucher []Voucher `json:"data"`
}
