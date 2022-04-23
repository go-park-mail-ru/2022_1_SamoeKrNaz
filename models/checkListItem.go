package models

type CheckListItem struct {
	IdClIt      uint   `json:"idclit" gorm:"primaryKey"`
	Description string `json:"title" gorm:"not null"`
	IsReady     bool   `json:"isready" gorm:"not null;default:false"`
}
