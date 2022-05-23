package models

//easyjson:json
type Lists []List

//easyjson:json
type List struct {
	IdL      uint   `json:"idl" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"not null"`
	Position uint   `json:"position" gorm:"not null"`
	IdB      uint   `json:"idb" gorm:"not null;"`
	Tasks    []Task `gorm:"foreignKey:IdL;constraint:OnDelete:CASCADE;"`
}
