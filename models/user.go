package models

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var UserID uint = 4

var UserList = []User{
	{1, "cucumber_two_two", "ya_lublu_kotikov"},
	{2, "my_friendISGood", "ya_lublu1213141"},
	{3, "xz_xz", "sobaki_toze_norm"},
}
