package usecases

import (
	"PLANEXA_backend/models"
)

func GetBoards(userId uint) ([]models.Board, error) {
	// достаю из БД доски по userId
	var err error // обработка ошибки из бд
	return models.BoardList, err
}
