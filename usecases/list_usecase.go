package usecases

import "PLANEXA_backend/models"

func GetLists(boardId uint, userId uint) ([]models.List, error) {
	// достаю все списки тасков из БД по айди доски
	var err error
	return []models.List{}, err
}

func GetSingleList(listId uint, userId uint) (models.List, error) {
	// доставю список из бд
	var err error
	return models.List{}, err
}

func CreateList(list models.List, userId uint) (uint, error) {
	// создаю список в бд, получаю айди листа
	var err error
	return 0, err
}

func RefactorList(list models.List, userId uint, boardId uint) error {
	// проверяю может ли юзер редачить
	// вношу изменения в бд
	var err error
	return err
}

func DeleteList(listId uint, userId uint) error {
	// проверяю есть ли такой лист и может ли юзер удалить его
	// удаляю список
	var err error
	return err
}
