package models

//easyjson:json
type Users []User

//easyjson:json
type User struct {
	IdU       uint      `json:"idu" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"not null;unique;index"`
	Password  string    `json:"password" gorm:"not null;"`
	ImgAvatar string    `json:"img_avatar"`
	Boards    []Board   `gorm:"many2many:users_boards;"`
	Tasks     []Task    `gorm:"many2many:users_tasks;constraint:OnDelete:CASCADE;"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:IdU;constraint:OnDelete:CASCADE;"`
}
