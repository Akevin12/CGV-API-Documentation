package models

type ReviewForm struct {
	TicketID int     `json:"ticket_id,omitempty"`
	Rating   float32 `json:"rating"`
	Content  string  `json:"content"`
	Date     string  `json:"date"`
}

type ReviewFormBody struct {
	ReviewForm ReviewForm `json:"data"`
}
