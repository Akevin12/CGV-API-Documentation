package models

type Ticket struct {
	ID       int            `json:"id"`
	Seat     Seat           `json:"seat"`
	Schedule ScheduleTicket `json:"schedule"`
	Payment  Payment        `json:"payment"`
}

type TicketC struct {
	ID       int            `json:"id"`
	Seat     Seat           `json:"seat"`
	Schedule ScheduleTicket `json:"schedule"`
	PaymentC PaymentC       `json:"payment"`
}

type TicketResponse struct {
	Response
	Ticket Ticket `json:"data"`
}
type TicketsResponse struct {
	Response
	Tickets []Ticket `json:"data"`
}

type TicketResponseC struct {
	Response
	TicketC TicketC `json:"data"`
}
type TicketsResponseC struct {
	Response
	TicketsC []TicketC `json:"data"`
}
