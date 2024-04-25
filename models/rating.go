package models

type ReviewSummary struct {
	Rating      float32 `json:"Rating"`
	TotalReview int     `json:"TotalReview"`
}
