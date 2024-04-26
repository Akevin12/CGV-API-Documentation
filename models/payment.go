package models

type Payment struct {
	ID     int           `json:"id"`
	Amount float64       `json:"amount"`
	Status PaymentStatus `json:"status"`
}

type PaymentC struct {
	ID     int            `json:"id"`
	Amount float64        `json:"amount"`
	Status PaymentStatusC `json:"status"`
}

type PaymentStatus string
type PaymentStatusC string

const (
	Pending  PaymentStatus  = "pending"
	Complete PaymentStatusC = "completed"
)

type PaymentResponse struct {
	Response
	Payment Payment `json:"data"`
}
type PaymentsResponse struct {
	Response
	Payments []Payment `json:"data"`
}
