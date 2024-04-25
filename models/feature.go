package models

type Feature struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	PicturePath string `json:"picture_path"`
	Link        string `json:"link"`
}

type Features struct {
	Features []Feature `json:"data"`
}

type FeaturesResponse struct {
	Response
	Features
}
