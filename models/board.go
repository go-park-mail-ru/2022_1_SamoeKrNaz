package models

type Board struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateCreated string `json:"date"`
}

var BoardList = []Board{
	{1, "Trello_project", "Today we wanna do something like trello", "22.02.2022"},
	{2, "Homework", "Lets talk about homework", "25.02.2022"},
	{3, "Math", "Let's do math", "10.02.2022"},
}
