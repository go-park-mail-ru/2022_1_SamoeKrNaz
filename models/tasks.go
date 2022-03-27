package models

import "time"

type Task struct {
	IdT         uint      `json:"idt"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Position    uint      `json:"position" gorm:"unique"`
	DateCreated time.Time `json:"dateCreated"`
	IdL         uint
	IdB         uint
}

var TaskList = []Task{
	{1, "РК № 1. СУБД", "4 марта 2022 года", 0, time.Now(), 0, 0},
	{2, "РК № 1. Фронт", "9 марта 2022 года", 0, time.Now(), 0, 0},
	{3, "РК № 1. Go", "10 марта 2022 года", 0, time.Now(), 0, 0},
}
