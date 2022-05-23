package models

//easyjson:json
type CheckLists []CheckList

//easyjson:json
type CheckList struct {
	IdCl           uint            `json:"id_cl" gorm:"primaryKey"`
	Title          string          `json:"title" gorm:"not null"`
	IdT            uint            `json:"id_t" gorm:"foreignKey:idCL"`
	CheckListItems []CheckListItem `gorm:"foreignKey:IdCl;constraint:OnDelete:CASCADE;"`
}
