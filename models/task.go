package models

import "time"

type Task struct {
	IdT         uint         `json:"idt" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description"`
	Position    uint         `json:"position" gorm:"not null"`
	DateCreated string       `json:"dateCreated"`
	IdL         uint         `gorm:"not null;"`
	IdB         uint         `gorm:"not null;"`
	DateToOrder time.Time    `gorm:"not null;"`
	Deadline    string       `json:"deadline"`
	IdU         uint         `gorm:"not null;"`
	IsReady     bool         `json:"is_ready" gorm:"not null;"`
	IsImportant bool         `json:"is_important" gorm:"not null;"`
	Link        string       `json:"link" gorm:"not null;"`
	CheckLists  []CheckList  `json:"checkList" gorm:"foreignKey:IdT;constraint:OnDelete:CASCADE;"`
	Comments    []Comment    `json:"comment" gorm:"foreignKey:IdT;constraint:OnDelete:CASCADE;"`
	Users       []User       `json:"append_users" gorm:"many2many:users_tasks;constraint:OnDelete:CASCADE;"`
	Attachments []Attachment `json:"attachments" gorm:"foreignKey:IdT;constraint:OnDelete:CASCADE;"`
}
