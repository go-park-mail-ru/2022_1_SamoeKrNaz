package models

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var UserID uint = 4

var UserList = []User{
	{1, "user1", "pass1"},
	{2, "user2", "pass2"},
	{3, "user3", "pass3"},
}
