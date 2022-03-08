package models

type Session struct {
	UserId      uint   `json:"user_id"`
	CookieValue string `json:"cookie_value"`
}

var SessionList = []Session{
	{1, "session1"},
	{2, "session2"},
	{3, "session3"},
}
