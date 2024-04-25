package models

type Review struct {
	Username    string  `json:"username"`
	PicturePath string  `json:"picture_path"`
	Rating      float32 `json:"Rating"`
	Content     string  `json:"content"`
	Date        string  `json:"date"`
}

type Reviews struct {
	Reviews []Review `json:"data"`
}

type ReviewsResponse struct {
	Response
	Reviews
}
