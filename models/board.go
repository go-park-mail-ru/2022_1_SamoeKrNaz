package models

type Board struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateCreated string `json:"date"`
}

var BoardList = []Board{
	{1, "Trello_project", "Today we wanna do something like trello", "22.02.2022"},
	{2, "I love kotikov", "We love cats and we are happy", "25.02.2022"},
	{3, "Cucumbers", "The best food is cucumber", "10.02.2022"},
}
