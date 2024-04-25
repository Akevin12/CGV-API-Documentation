package models

type Points struct {
	Total  int    `json:"total"`
	Date   string `json:"date"`
	Expire string `json:"expire"`
}

type PointsResponse struct {
	Response
	Points Points `json:"data"`
}
