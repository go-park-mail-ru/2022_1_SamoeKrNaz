package models

import "time"

type Board struct {
	IdB         uint   `json:"idb" gorm:"primaryKey;auto_increment"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	ImgDesk     string `json:"img_desk"`
	DateCreated string `json:"date"`
	IdU         uint   `json:"idp" gorm:"foreignKey:IdB"`
	Users       []User `gorm:"many2many:users_boards"`
	Lists       []List `gorm:"foreignKey:IdB"`
	Tasks       []Task `gorm:"foreignKey:IdB"`
}

var BoardList = []Board{
	{1, "Заголовок доски", "", "Сегодня мы решили сделать что-то типо trello", time.Now().String(), 1, nil, nil, nil},
	{2, "Домашка", "", "Тут я буду строить планы по выполнению домашнего задания для школы.", time.Now().String(), 2, nil, nil, nil},
	{3, "Математика", "", "Давайте позанимаемся математикой", time.Now().String(), 3, nil, nil, nil},
	{4, "Домашние дела", "", "Тут собраны все домашние дела, которые мне надо выполнить", time.Now().String(), 4, nil, nil, nil},
	{5, "Проект", "", "Все таски проекта собраны тут, для того, чтобы мы могли эффективно вести разработку вместе. Это точно удобно!", time.Now().String(), 5, nil, nil, nil},
}

type BoardAndTasks struct {
	Boards []Board `json:"boards"`
	Tasks  []Task  `json:"tasks"`
}

var TasksAndBoards = BoardAndTasks{Boards: BoardList, Tasks: TaskList}
