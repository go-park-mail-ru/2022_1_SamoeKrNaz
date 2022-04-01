package usecases

import (
	"PLANEXA_backend/models"
)

func GetBoards(userId uint) ([]models.Board, error) {
	// достаю из БД доски по userId
	var err error // обработка ошибки из бд
	return models.BoardList, err
}

func GetSingleBoard(boardId uint, userId uint) (models.Board, error) {
	//проверить может ли юзер смотреть эту доску
	// вызываю из бд получение доски
	var err error
	// обрабатываю ошибку
	return models.Board{}, err
}

func CreateBoard(userId uint, board models.Board) error {
	// добавляю в бд такую доску с привязкой к данному юзеру
	var err error // обрабатываю ошибку из бд
	return err
}

func RefactorBoard(userId uint, board models.Board) error {
	// проверяю есть ли доска с таким айди и может ли юзер её редачить
	//вызываю репозиторий дляобновления доски
	var err error // обрабатвываю ошибку
	return err
}

func DeleteBoard(boardId uint, userId uint) error {
	// проверяю есть ли такая доска и может ли юзер редачить её
	// удаляю из бд
	var err error
	return err
}
