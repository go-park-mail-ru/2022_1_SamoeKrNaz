package models

//easyjson:json
type Session struct {
	UserId      uint   `json:"user_id"`
	CookieValue string `json:"cookie_value"`
}
