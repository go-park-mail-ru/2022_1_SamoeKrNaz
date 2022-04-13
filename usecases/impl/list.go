package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories/impl"
	"PLANEXA_backend/usecases"
	"github.com/microcosm-cc/bluemonday"
)

type ListUseCaseImpl struct {
	repList  *impl.ListRepository
	repBoard *impl.BoardRepository
}

func MakeListUsecase(repList_ *impl.ListRepository, repBoard_ *impl.BoardRepository) usecases.ListUseCase {
	return &ListUseCaseImpl{repList: repList_, repBoard: repBoard_}
}

func (listUseCase *ListUseCaseImpl) GetLists(boardId uint, userId uint) ([]models.List, error) {
	// достаю все списки тасков из БД по айди доски
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	lists, err := listUseCase.repList.GetLists(boardId)
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
	} else if isAccess == false {
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

func (listUseCase *ListUseCaseImpl) CreateList(list models.List, boardId uint, userId uint) (uint, error) {
	// создаю список в бд, получаю айди листа
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return 0, err
	} else if isAccess == false {
		return 0, customErrors.ErrNoAccess
	}
	listId, err := listUseCase.repList.Create(&list, boardId)
	return listId, err
}

func (listUseCase *ListUseCaseImpl) RefactorList(list models.List, userId uint, boardId uint) error {
	// проверяю может ли юзер редачить
	// вношу изменения в бд
	isAccess, err := listUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return err
	} else if isAccess == false {
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
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	_, err = listUseCase.repList.GetById(listId)
	if err != nil {
		return err
	}

	err = listUseCase.repList.Delete(listId)
	return err
}
