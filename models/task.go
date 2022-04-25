package models

type Task struct {
	IdT         uint        `json:"idt" gorm:"primaryKey"`
	Title       string      `json:"title" gorm:"not null"`
	Description string      `json:"description"`
	Position    uint        `json:"position" gorm:"not null""`
	DateCreated string      `json:"dateCreated"`
	IdL         uint        `gorm:"not null;"`
	IdB         uint        `gorm:"not null;"`
	CheckLists  []CheckList `json:"checkList" gorm:"foreignKey:IdT;constraint:OnDelete:CASCADE;"`
	Comments    []Comment   `json:"comment" gorm:"foreignKey:IdT;constraint:OnDelete:CASCADE;"`
}
