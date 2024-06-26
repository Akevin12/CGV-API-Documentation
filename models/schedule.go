package models

import "time"

type Schedule struct {
	ID       int       `json:"id"`
	Price    float64   `json:"price"`
	Showtime time.Time `json:"showtime"`
	Movie    *Movie    `json:"movie,omitempty"`
	Branch   Branch    `json:"branch"`
	Seats    *[]Seat   `json:"seats,omitempty"`
}

type ScheduleTicket struct {
	ID       int        `json:"id"`
	Price    *float64   `json:"price,omitempty"`
	Showtime *time.Time `json:"showtime,omitempty"`
	Movie    *Movie     `json:"movie,omitempty"`
	Branch   *Branch    `json:"branch,omitempty"`
	Seat     *Seat      `json:"seat,omitempty"`
}

type SchedulesResponse struct {
	Response
	Schedules []Schedule `json:"data"`
}
type ScheduleResponse struct {
	Response
	Schedule Schedule `json:"data"`
}
