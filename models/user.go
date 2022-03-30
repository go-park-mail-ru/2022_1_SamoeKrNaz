package models

type User struct {
	IdU       uint    `json:"idu" gorm:"primaryKey;auto_increment"`
	Username  string  `json:"username" gorm:"not null;unique;index"`
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
