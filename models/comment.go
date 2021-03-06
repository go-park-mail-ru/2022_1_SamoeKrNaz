package models

//easyjson:json
type Comments []Comment

//easyjson:json
type Comment struct {
	IdCm        uint   `json:"idcm" gorm:"primaryKey"`
	Text        string `json:"title" gorm:"not null"`
	DateCreated string `json:"date"`
	IdT         uint   `json:"idt" gorm:"not null"`
	IdU         uint   `json:"idu" gorm:"not null"`
	User        User   `json:"user" gorm:"-"`
}
