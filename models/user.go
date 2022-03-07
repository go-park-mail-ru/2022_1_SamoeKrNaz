package models

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var UserID uint = 4

var UserList = []User{
	{1, "palantina14", "bdazglweq21"},
	{2, "Lopp", "1labwaf2"},
	{3, "paperThing11", "gedab1gawf"},
}
