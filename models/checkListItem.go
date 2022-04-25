package models

type CheckListItem struct {
	IdClIt      uint   `json:"id_clit" gorm:"primaryKey"`
	Description string `json:"title" gorm:"not null"`
	IdCl        uint   `json:"id_cl" gorm:"not null"`
	IdT         uint   `json:"id_t" gorm:"not null"`
	IsReady     bool   `json:"isready" gorm:"not null;default:false"`
}
