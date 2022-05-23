package models

//easyjson:json
type Boards []Board

//easyjson:json
type Board struct {
	IdB         uint   `json:"idb" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	ImgDesk     string `json:"img_desk"`
	DateCreated string `json:"date"`
	IdU         uint   `json:"idu" gorm:"foreignKey:IdB;"`
	Link        string `json:"link" gorm:"not null;"`
	Users       []User `gorm:"many2many:users_boards"`
	Lists       []List `gorm:"foreignKey:IdB;constraint:OnDelete:CASCADE;"`
	Tasks       []Task `gorm:"foreignKey:IdB;constraint:OnDelete:CASCADE;"`
}
