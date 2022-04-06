package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
	"time"
)

type BoardUseCaseImpl struct {
	rep *repositories.BoardRepository
}

func MakeBoardUsecase(rep_ *repositories.BoardRepository) usecases.BoardUseCase {
	return &BoardUseCaseImpl{rep: rep_}
}

func (boardUseCase *BoardUseCaseImpl) GetBoards(userId uint) ([]models.Board, error) {
	// достаю из БД доски по userId
	boards, err := boardUseCase.rep.GetUserBoards(userId)
	if err != nil {
		return nil, err
	}
	return boards, nil
}

func (boardUseCase *BoardUseCaseImpl) GetSingleBoard(boardId uint, userId uint) (models.Board, error) {
	//проверить может ли юзер смотреть эту доску
	// вызываю из бд получение доски
	board, err := boardUseCase.rep.GetById(boardId)
	if err != nil {
		return models.Board{}, err
	}

	if board.IdU != userId {
		return models.Board{}, customErrors.ErrUserHasntBoards
	}
	// обрабатываю ошибку
	return *board, nil
}

func (boardUseCase *BoardUseCaseImpl) CreateBoard(userId uint, board models.Board) error {
	// добавляю в бд такую доску с привязкой к данному юзеру
	board.DateCreated = time.Now().Format(time.RFC850)
	board.IdU = userId
	err := boardUseCase.rep.Create(board)
	return err
}

func (boardUseCase *BoardUseCaseImpl) RefactorBoard(userId uint, board models.Board) error {
	// проверяю есть ли доска с таким айди и может ли юзер её редачить
	//вызываю репозиторий дляобновления доски
	if board.IdU != userId {
		return customErrors.ErrUnauthorized
	}
	err := boardUseCase.rep.Update(board)
	return err
}

func (boardUseCase *BoardUseCaseImpl) DeleteBoard(boardId uint, userId uint) error {
	// проверяю есть ли такая доска и может ли юзер редачить её
	// удаляю из бд
	board, err := boardUseCase.rep.GetById(boardId)
	if err != nil {
		return customErrors.ErrBoardNotFound
	}
	if board.IdU != userId {
		return customErrors.ErrAccess
	}

	err = boardUseCase.rep.Delete(boardId)
	return err
}
