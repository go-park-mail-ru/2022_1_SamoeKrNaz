package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories/impl"
	"PLANEXA_backend/usecases"
	"github.com/microcosm-cc/bluemonday"
	"time"
)

type BoardUseCaseImpl struct {
	rep *impl.BoardRepository
}

func MakeBoardUsecase(rep_ *impl.BoardRepository) usecases.BoardUseCase {
	return &BoardUseCaseImpl{rep: rep_}
}

func (boardUseCase *BoardUseCaseImpl) GetBoards(userId uint) ([]models.Board, error) {
	// достаю из БД доски по userId
	boards, err := boardUseCase.rep.GetUserBoards(userId)
	if err != nil {
		return nil, err
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, board := range boards {
		board.DateCreated = sanitizer.Sanitize(board.DateCreated)
		board.Title = sanitizer.Sanitize(board.Title)
		board.Description = sanitizer.Sanitize(board.Description)
		board.ImgDesk = sanitizer.Sanitize(board.ImgDesk)
	}
	return boards, nil
}

func (boardUseCase *BoardUseCaseImpl) GetSingleBoard(boardId uint, userId uint) (models.Board, error) {
	//проверить может ли юзер смотреть эту доску
	// вызываю из бд получение доски
	isAccess, err := boardUseCase.rep.IsAccessToBoard(userId, boardId)
	if err != nil {
		return models.Board{}, err
	} else if isAccess == false {
		return models.Board{}, customErrors.ErrNoAccess
	}

	board, err := boardUseCase.rep.GetById(boardId)
	if err != nil {
		return models.Board{}, err
	}
	sanitizer := bluemonday.UGCPolicy()
	board.DateCreated = sanitizer.Sanitize(board.DateCreated)
	board.Title = sanitizer.Sanitize(board.Title)
	board.Description = sanitizer.Sanitize(board.Description)
	board.ImgDesk = sanitizer.Sanitize(board.ImgDesk)
	return *board, nil
}

func (boardUseCase *BoardUseCaseImpl) CreateBoard(userId uint, board models.Board) (uint, error) {
	// добавляю в бд такую доску с привязкой к данному юзеру
	board.DateCreated = time.Now().Format(time.RFC850)
	board.IdU = userId
	boardId, err := boardUseCase.rep.Create(&board)
	if err != nil {
		return 0, err
	}
	err = boardUseCase.rep.AppendUser(&board)
	if err != nil {
		return 0, err
	}
	return boardId, nil
}

func (boardUseCase *BoardUseCaseImpl) RefactorBoard(userId uint, board models.Board) error {
	// проверяю есть ли доска с таким айди и может ли юзер её редачить
	//вызываю репозиторий для обновления доски
	isAccess, err := boardUseCase.rep.IsAccessToBoard(userId, board.IdB)
	if err != nil {
		return err
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	err = boardUseCase.rep.Update(board)
	return err
}

func (boardUseCase *BoardUseCaseImpl) DeleteBoard(boardId uint, userId uint) error {
	// проверяю есть ли такая доска и может ли юзер редачить её
	// удаляю из бд
	isAccess, err := boardUseCase.rep.IsAccessToBoard(userId, boardId)
	if err != nil {
		return err
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	err = boardUseCase.rep.Delete(boardId)
	return err
}
