package impl

import (
	customErrors "PLANEXA_backend/errors"
	repositories "PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/microcosm-cc/bluemonday"
)

type ListUseCaseImpl struct {
	repList  repositories.ListRepository
	repBoard repositories.BoardRepository
}

func MakeListUsecase(repList_ repositories.ListRepository, repBoard_ repositories.BoardRepository) usecases.ListUseCase {
	return &ListUseCaseImpl{repList: repList_, repBoard: repBoard_}
}

func (listUseCase *ListUseCaseImpl) GetLists(boardId uint, userId uint) ([]models.List, error) {
	// достаю все списки тасков из БД по айди доски
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	lists, err := listUseCase.repBoard.GetLists(boardId)
	if err != nil {
		return nil, err
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, list := range lists {
		list.Title = sanitizer.Sanitize(list.Title)
	}
	return lists, err
}

func (listUseCase *ListUseCaseImpl) GetSingleList(listId uint, userId uint) (models.List, error) {
	// доставю список из бд
	board, err := listUseCase.repList.GetBoard(listId)
	if err != nil {
		return models.List{}, err
	}
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, board.IdB)
	if err != nil {
		return models.List{}, err
	} else if !isAccess {
		return models.List{}, customErrors.ErrNoAccess
	}
	list, err := listUseCase.repList.GetById(listId)
	if err != nil {
		return models.List{}, err
	}
	sanitizer := bluemonday.UGCPolicy()
	list.Title = sanitizer.Sanitize(list.Title)
	return *list, err
}

func (listUseCase *ListUseCaseImpl) CreateList(list models.List, boardId uint, userId uint) (*models.List, error) {
	// создаю список в бд, получаю айди листа
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	listId, err := listUseCase.repList.Create(&list, boardId)
	if err != nil {
		return nil, err
	}
	createdList, err := listUseCase.repList.GetById(listId)
	return createdList, err
}

func (listUseCase *ListUseCaseImpl) RefactorList(list models.List, userId uint, boardId uint) error {
	// проверяю может ли юзер редачить
	// вношу изменения в бд
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = listUseCase.repList.Update(list)
	return err
}

func (listUseCase *ListUseCaseImpl) DeleteList(listId uint, userId uint) error {
	// проверяю есть ли такой лист и может ли юзер удалить его
	// удаляю список
	board, err := listUseCase.repList.GetBoard(listId)
	if err != nil {
		return err
	}
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, board.IdB)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	_, err = listUseCase.repList.GetById(listId)
	if err != nil {
		return err
	}
	err = listUseCase.repList.Delete(listId)
	return err
}
