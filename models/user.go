package models

// User struct data
type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Paternal string `json:"paternal"`
	Maternal string `json:"maternal"`
	Email    string `json:"email"`
}
