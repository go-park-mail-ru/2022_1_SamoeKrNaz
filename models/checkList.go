package models

type CheckList struct {
	IdCl           uint            `json:"idcl" gorm:"primaryKey"`
	Title          string          `json:"title" gorm:"not null"`
	IdT            uint            `json:"idt" gorm:"foreignKey:idCL"`
	CheckListItems []CheckListItem `gorm:"foreignKey:IdCl;constraint:OnDelete:CASCADE;"`
}
