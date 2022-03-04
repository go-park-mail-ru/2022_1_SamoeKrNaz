package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	SessionId int    `json:"session_Id"`
	CookieId  string `json:"cookie_id"`
}

type Board struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Date        string `json:"date"`
}

var UserID int
var UserList = []User{
	{1, "user1", "pass1"},
	{2, "user2", "pass2"},
	{3, "user3", "pass3"},
}
var SessionList []Session
var BoardList = []Board{
	{1, "board1", "descr1", "img/img1", "22.02.2022"},
	{2, "board2", "descr2", "img/img2", "22.02.2023"},
	{3, "board3", "descr3", "img/img3", "22.02.2024"},
}
