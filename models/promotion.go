package models

type Promotion struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	PicturePath string `json:"picture_path"`
	Link        string `json:"link"`
	Date        string `json:"date"`
}

type Promotions struct {
	Promotions []Promotion `json:"data"`
}

type PromotionsResponse struct {
	Response
	Promotions
}
