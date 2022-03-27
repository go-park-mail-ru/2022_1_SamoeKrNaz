package models

type List struct {
	IdL       uint   `json:"idl" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null"`
	IdB       uint   `json:"idb"`
	ImgAvatar string `json:"img_avatar"`
	Tasks     []Task `gorm:"foreignKey:IdL"`
}
