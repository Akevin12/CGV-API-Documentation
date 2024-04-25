package models

import "time"

type Branch struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BranchDetail struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	MallAccess  string     `json:"mall access"`
	ParkingArea string     `json:"parking area`
	Operational *time.Time `json: "operational hours"`
}

type BranchesResponse struct {
	Response
	Branches []Branch `json:"data"`
}
type BranchResponse struct {
	Response
	Branch Branch `json:"data"`
}
type BranchDetailResponse struct {
	Response
	BranchDetail BranchDetail `json:"data"`
}
