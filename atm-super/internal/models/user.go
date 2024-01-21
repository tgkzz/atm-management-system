package models

type User struct {
	Id          string
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Description string `json:"description,omitempty"`
}