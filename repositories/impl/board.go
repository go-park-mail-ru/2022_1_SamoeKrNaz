package impl

import (
	"PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"gorm.io/gorm"
)

type BoardRepositoryImpl struct {
	db *gorm.DB
}

func MakeBoardRepository(db *gorm.DB) repositories.BoardRepository {
	return &BoardRepositoryImpl{db: db}
}

func (boardRepository *BoardRepositoryImpl) Create(board *models.Board) (uint, error) {
	err := boardRepository.db.Create(board).Error
	return board.IdB, err
}

func (boardRepository *BoardRepositoryImpl) AppendUser(board *models.Board) error {
	err := boardRepository.db.Model(&models.User{IdU: board.IdU}).Association("Boards").Append(board)
	return err
}

func (boardRepository *BoardRepositoryImpl) GetLists(IdB uint) ([]models.List, error) {
	lists := new([]models.List)
	result := boardRepository.db.Where("id_b = ?", IdB).Order("position").Find(lists)
	return *lists, result.Error
}

func (boardRepository *BoardRepositoryImpl) Update(board models.Board) error {
	// возьмем из бд текущую запись по айдишнику
	currentData, err := boardRepository.GetById(board.IdB)
	// обработка ошибки при взятии
	if err != nil {
		return err
	}
	// ищем, какое поле поменялось
	if currentData.Title != board.Title && board.Title != "" {
		currentData.Title = board.Title
	}
	if currentData.Description != board.Description && board.Description != "" {
		currentData.Description = board.Description
	}
	//сохраняем новую структуру
	return boardRepository.db.Save(currentData).Error
}

func (boardRepository *BoardRepositoryImpl) Delete(IdB uint) error {
	err := boardRepository.db.Model(&models.Board{IdB: IdB}).Association("Users").Clear()
	if err != nil {
		return err
	}
	return boardRepository.db.Delete(&models.Board{}, IdB).Error
}

func (boardRepository *BoardRepositoryImpl) GetUserBoards(IdU uint) ([]models.Board, error) {
	boards := new([]models.Board)
	err := boardRepository.db.Model(&models.User{IdU: IdU}).Association("Boards").Find(boards)
	if err != nil {
		return nil, err
	}
	return *boards, nil
}

func (boardRepository *BoardRepositoryImpl) GetById(IdB uint) (*models.Board, error) {
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

func (boardRepository *BoardRepositoryImpl) IsAccessToBoard(IdU uint, IdB uint) (bool, error) {
	board := new(models.Board)
	err := boardRepository.db.Model(&models.User{IdU: IdU}).Where("id_b = ?", IdB).Association("Boards").Find(board)
	if err != nil {
		return false, err
	} else if board == nil {
		return false, nil
	}
	return true, nil
}
