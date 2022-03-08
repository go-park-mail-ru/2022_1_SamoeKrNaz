package models

type Task struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

var TaskList = []Task{
	{1, "РК № 1. СУБД", "4 марта 2022 года"},
	{2, "РК № 1. Фронт", "9 марта 2022 года"},
	{3, "РК № 1. Go", "10 марта 2022 года"},
}
