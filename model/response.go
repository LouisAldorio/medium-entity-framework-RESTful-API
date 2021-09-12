package model

type Response struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Error  string      `json:"message"`
	Data   interface{} `json:"data"`
}
