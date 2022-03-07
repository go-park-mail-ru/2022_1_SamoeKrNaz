package models

type Board struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateCreated string `json:"date"`
}

var BoardList = []Board{
	{1, "Trello_project", "Today we wanna do something like trello", "22.02.2022"},
	{2, "Food talks", "Lets talk about food", "25.02.2022"},
	{3, "Cucumber", "The best food is cucumber", "10.02.2022"},
}
