package repositories

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
	"time"
)

type BoardRepository struct {
	db *gorm.DB
}

func (boardRepository *BoardRepository) MakeRepository(db *gorm.DB) *BoardRepository {
	return &BoardRepository{db: db}
}

func (boardRepository *BoardRepository) Create(board *models.Board, IdU uint) error {
	// TODO: вынести две нижние строки в юзкейс
	board.DateCreated = time.Now().Format(time.RFC850)
	board.IdU = IdU
	return boardRepository.db.Create(board).Error
}

func (boardRepository *BoardRepository) Update(board *models.Board) error {
	// возьмем из бд текущую запись по айдишнику
	currentData, err := boardRepository.GetById(board.IdB)
	// обработка ошибки при взятии
	if err != nil {
		return err
	}
	// ищем, какое поле поменялось
	if currentData.Title != board.Title {
		currentData.Title = board.Title
	}
	if currentData.Description != board.Description {
		currentData.Description = board.Description
	}
	//сохраняем новую структуру
	return boardRepository.db.Save(currentData).Error
}

func (boardRepository *BoardRepository) Delete(IdB uint) error {
	return boardRepository.db.Delete(&models.Board{}, IdB).Error
}

func (boardRepository *BoardRepository) GetUserBoards(IdU uint) (*[]models.Board, error) {
	boards := new([]models.Board)
	err := boardRepository.db.Model(&models.User{IdU: IdU}).Association("Boards").Find(boards)
	if err != nil {
		return nil, err
	}
	return boards, nil
}

func (boardRepository *BoardRepository) GetLists(IdB uint) (*[]models.List, error) {
	lists := new([]models.List)
	result := boardRepository.db.Where("id_b = ?", IdB).Find(lists)
	return lists, result.Error
}

func (boardRepository *BoardRepository) GetTasks(IdB uint) (*[]models.Task, error) {
	tasks := new([]models.Task)
	result := boardRepository.db.Where("id_b = ?", IdB).Find(tasks)
	return tasks, result.Error
}

func (boardRepository *BoardRepository) GetById(IdB uint) (*models.Board, error) {
	// указатель на структуру, которую вернем
	board := new(models.Board)
	result := boardRepository.db.Find(board, IdB)
	// если выборка в 0 строк, то такой доски нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrBoardNotFound
	} else if result.Error != nil {
		// если произошла ошибка при выборке
		return nil, result.Error
	}
	// иначе вернем доску
	return board, nil
}
