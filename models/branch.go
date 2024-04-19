package models

type Branch struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BranchesResponse struct {
	Response
	Branches []Branch `json:"data"`
}
type BranchResponse struct {
	Response
	Branch Branch `json:"data"`
}
