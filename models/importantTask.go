package models

type ImportantTask struct {
	IdU uint `gorm:"not null;"`
	IdT uint `gorm:"not null;"`
	IdB uint `gorm:"not null;"`
}
