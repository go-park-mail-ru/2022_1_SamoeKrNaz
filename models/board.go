package models

import (
	"time"
)

type Board struct {
	IdB         uint      `json:"idb" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	ImgDesk     string    `json:"img_desk"`
	DateCreated time.Time `json:"date"`
	IdP         uint      `json:"idp" gorm:"foreignKey:IdP"`
	Users       []User    `gorm:"many2many:users_boards"`
	Tasks       []Task    `gorm:"foreignKey:IdB"`
}

var BoardList = []Board{
	{1, "Заголовок доски", "", "Сегодня мы решили сделать что-то типо trello", time.Now(), 1, nil, nil},
	{2, "Домашка", "", "Тут я буду строить планы по выполнению домашнего задания для школы.", time.Now(), 2, nil, nil},
	{3, "Математика", "", "Давайте позанимаемся математикой", time.Now(), 3, nil, nil},
	{4, "Домашние дела", "", "Тут собраны все домашние дела, которые мне надо выполнить", time.Now(), 4, nil, nil},
	{5, "Проект", "", "Все таски проекта собраны тут, для того, чтобы мы могли эффективно вести разработку вместе. Это точно удобно!", time.Now(), 5, nil, nil},
}

type BoardAndTasks struct {
	Boards []Board `json:"boards"`
	Tasks  []Task  `json:"tasks"`
}

var TasksAndBoards = BoardAndTasks{Boards: BoardList, Tasks: TaskList}
