package impl

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
)

type ListUseCaseImpl struct {
	rep *repositories.ListRepository
}

func MakeListUsecase(rep_ *repositories.ListRepository) usecases.ListUseCase {
	return &ListUseCaseImpl{rep: rep_}
}

func (listUseCase *ListUseCaseImpl) GetLists(boardId uint, userId uint) ([]models.List, error) {
	// достаю все списки тасков из БД по айди доски
	lists, err := listUseCase.rep.GetLists(boardId)
	return lists, err
}

func (listUseCase *ListUseCaseImpl) GetSingleList(listId uint) (models.List, error) {
	// доставю список из бд
	list, err := listUseCase.rep.GetById(listId)
	return *list, err
}

func (listUseCase *ListUseCaseImpl) CreateList(list models.List, boardId uint) (uint, error) {
	// создаю список в бд, получаю айди листа
	err := listUseCase.rep.Create(list, boardId)

	return 0, err
}

func (listUseCase *ListUseCaseImpl) RefactorList(list models.List, userId uint, boardId uint) error {
	// проверяю может ли юзер редачить
	// вношу изменения в бд
	err := listUseCase.rep.Update(list)
	return err
}

func (listUseCase *ListUseCaseImpl) DeleteList(listId uint, userId uint) error {
	// проверяю есть ли такой лист и может ли юзер удалить его
	// удаляю список
	_, err := listUseCase.rep.GetById(listId)
	if err != nil {
		return err
	}

	err = listUseCase.rep.Delete(listId)
	return err
}
