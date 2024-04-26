package models

type Voucher struct {
	IDVoucher   int    `json:"idvoucher"`
	VoucherName string `json:"vouchername"`
	Description string `json:"description"`
	ExpiredDate string `json:"expireddate`
}

type VoucherResponse struct {
	Response
	Voucher []Voucher `json:"data"`
}
