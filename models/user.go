package models

type User struct {
	ID           int    `json:"id"`
	EmailAddress string `json:"email_address" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Name         string `json:"name"`
}

type AuthUser struct {
	EmailAddress string `json:"email_address" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

func NameTableUser() string {
	return "users"
}
