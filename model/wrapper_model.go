package model

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	Expired string `json:"expired"`
	Token   string `json:"token"`
}
