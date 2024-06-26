package models

type Customer struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password *string `json:"password,omitempty"`
	Birthday string  `json:"birthday`
	Gender   string  `json:gender`
	Phone    string  `json:"phone"`
	Address  string  `json:"address"`
	City     string  `json:"city"`
	Cinema   Branch  `json:"cinema`
}

type CustomerResponse struct {
	Response
	Customer Customer `json:"data"`
}

type CustomerEdit struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:gender`
	Address string `json:"address"`
}

type Password struct {
	ID          int     `json:"id`
	CurrentPass *string `json:"current password,omitempty"`
	NewPass     *string `json:"new password,omitempty"`
	ConfrimPass *string `json:"confirm password,omitempty"`
}

type PassResponse struct {
	Response
	Password Password `json:"data"`
}
