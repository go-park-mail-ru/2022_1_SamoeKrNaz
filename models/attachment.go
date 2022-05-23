package models

//easyjson:json
type Attachments []Attachment

//easyjson:json
type Attachment struct {
	IdA         uint   `json:"id_a" gorm:"primaryKey"`
	DefaultName string `json:"default_name" gorm:"not null;"`
	SystemName  string `json:"system_name" gorm:"not null;"`
	IdT         uint   `json:"id_t" gorm:"not null;"`
}
