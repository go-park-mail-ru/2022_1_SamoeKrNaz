package models

type Comment struct {
	IdCm        uint   `json:"idcm" gorm:"primaryKey"`
	Text        string `json:"title" gorm:"not null"`
	DateCreated string `json:"date"`
	IdT         uint   `json:"idt" gorm:"not null"`
	User        User   `json:"user"`
}
