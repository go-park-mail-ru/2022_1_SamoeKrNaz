package usecases

import "PLANEXA_backend/models"

func GetTasks(listId uint, userId uint) ([]models.List, error) {
	// достаю все списки тасков из БД по айди доски
	var err error
	return []models.List{}, err
}

func GetSingleTask(listId uint, userId uint) (models.List, error) {
	// доставю список из бд
	var err error
	return models.List{}, err
}

func CreateTaks(list models.Task, userId uint) (uint, error) {
	// создаю список в бд, получаю айди листа
	var err error
	return 0, err
}

func RefactorTask(task models.Task, userId uint, listId uint) error {
	// проверяю может ли юзер редачить
	// вношу изменения в бд
	var err error
	return err
}

func DeleteTask(taskId uint, userId uint) error {
	// проверяю есть ли такой лист и может ли юзер удалить его
	// удаляю список
	var err error
	return err
}
