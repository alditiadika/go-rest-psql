package model

//ErrorModel struct
type ErrorModel struct {
	IsError bool        `json:"is_error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
