package model

type ResponseError struct {
	Title  string `json:"title"`
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}
