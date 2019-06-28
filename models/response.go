package models

// Response struct data
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []User `json:"data,omitempty"`
}
