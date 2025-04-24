package model

type WebResponse[T any] struct {
	Errors error `json:"error"`
	Data   T     `json:"data"`
	Status int   `json:"status_code"`
}
