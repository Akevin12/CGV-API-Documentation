package models

type News struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	PicturePath string `json:"picture_path"`
	Link        string `json:"link"`
	Date        string `json:"date"`
}

type NewsArray struct {
	ArrNews []News `json:"data"`
}

type NewsArrayResponse struct {
	Response
	NewsArray
}
