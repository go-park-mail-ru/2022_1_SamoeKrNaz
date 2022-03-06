package models

type Board struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateCreated string `json:"date"`
}

var BoardList = []Board{
	{1, "board1", "descr1", "22.02.2022"},
	{2, "board2", "descr2", "22.02.2023"},
	{3, "board3", "descr3", "22.02.2024"},
}
