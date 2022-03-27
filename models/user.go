package models

type User struct {
	IdP       uint    `json:"idp" gorm:"primaryKey"`
	Username  string  `json:"username" gorm:"not null;"`
	Password  string  `json:"password" gorm:"not null;"`
	ImgAvatar string  `json:"img_avatar"`
	Boards    []Board `gorm:"many2many:users_boards;"`
}

var UserID uint = 4

var UserList = []User{
	{1, "palantina14", "bdazglweq21", "", nil},
	{2, "Lopp", "1labwaf2", "", nil},
	{3, "paperThing11", "gedab1gawf", "", nil},
}
