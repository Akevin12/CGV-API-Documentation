package models

type Member struct {
	ID           int    `json:"id"`
	TypeMember   string `json:"type`
	CurrentLevel int    `json:"level`
	Points       string `json:"point`
	Spend        string `json:"spend`
	Benefit      string `json:"benefit"`
}

type MemberResponse struct {
	Response
	Member Member `json:"data"`
}
