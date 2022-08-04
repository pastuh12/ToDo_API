package models

type Session struct {
	ID           int    `json: "id"`
	UserID       int    `json: "userID"`
	RefreshToken string `json: "refreshTpoken"`
	ExpiresAt    int64  `json: "expiresAt"`
}
