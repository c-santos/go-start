package models

type Response struct {
	Message interface{} `json:"message"`
	Status  int    `json:"status"`
}
