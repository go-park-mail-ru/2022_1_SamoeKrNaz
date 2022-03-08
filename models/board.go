package models

type Board struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateCreated string `json:"date"`
}

var BoardList = []Board{
	{1, "Заголовок доски", "Сегодня мы решили сделать что-то типо trello", "22.02.2022"},
	{2, "Домашка", "Тут я буду строить планы по выполнению домашнего задания для школы.", "25.02.2022"},
	{3, "Математика", "Давайте позанимаемся математикой", "10.02.2022"},
	{4, "Домашние дела", "Тут собраны все домашние дела, которые мне надо выполнить", "10.02.2022"},
	{5, "Проект", "Все таски проекта собраны тут, для того, чтобы мы могли эффективно вести разработку вместе. Это точно удобно!", "10.02.2022"},
}

type BoardAndTasks struct {
	Boards []Board `json:"boards"`
	Tasks  []Task  `json:"tasks"`
}

var TasksAndBoards = BoardAndTasks{Boards: BoardList, Tasks: TaskList}
