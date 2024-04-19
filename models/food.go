package models

type Chart struct {
	ID       int    `json:"id"`
	Menu     string `json:"menu"`
	Quantity int    `json:"quantity`
	Price    string `json:"price"`
}

type ChartResponse struct {
	Response
	Chart Chart `json:"data"`
}
