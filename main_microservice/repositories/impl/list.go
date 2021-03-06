package impl

import (
	"PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
)

type ListRepositoryImpl struct {
	db *gorm.DB
}

func MakeListRepository(db *gorm.DB) repositories.ListRepository {
	return &ListRepositoryImpl{db: db}
}

func (listRepository *ListRepositoryImpl) Create(list *models.List, IdB uint) (uint, error) {
	list.IdB = IdB
	var currentPosition int64
	err := listRepository.db.Model(&models.List{}).Where("id_b = ?", list.IdB).Count(&currentPosition).Error
	if err != nil {
		return 0, err
	}
	list.Position = uint(currentPosition) + 1
	err = listRepository.db.Create(list).Error
	return list.IdL, err
}

func (listRepository *ListRepositoryImpl) Update(list models.List) error {
	currentData, err := listRepository.GetById(list.IdL)
	if err != nil {
		return err
	}
	if currentData.Title != list.Title && list.Title != "" {
		currentData.Title = list.Title
	}
	if currentData.Position != list.Position && list.Position != 0 {
		// если список переместили вниз
		if currentData.Position > list.Position {
			// допустим, что был список 1 2 3 4
			// решили, что четвертый список будет после первого
			// 1 4 2 3
			// значит, нужно все индексы после текущей позиции увеличить на 1
			err := listRepository.db.Model(&models.List{}).
				Where("id_b = ? AND position BETWEEN ? AND ?", currentData.IdB, list.Position, currentData.Position-1).
				UpdateColumn("position", gorm.Expr("position + 1")).Error
			if err != nil {
				return err
			}
			currentData.Position = list.Position
		} else { // если список переместили вверх
			// допустим, что был список 1 2 3 4
			// решили, что второй список будет после четвертого
			// 1 3 4 2
			// значит, нужно все индексы  с предыдущей позиции уменьшить на 1
			err := listRepository.db.Model(&models.List{}).
				Where("id_b = ? AND position BETWEEN ? AND ?", currentData.IdB, currentData.Position+1, list.Position).
				UpdateColumn("position", gorm.Expr("position - 1")).Error
			if err != nil {
				return err
			}
			currentData.Position = list.Position
		}
	}
	//сохраняем новую структуру
	return listRepository.db.Save(currentData).Error
}
func (listRepository *ListRepositoryImpl) Delete(IdL uint) error {
	// при удалении необходимо изменить позиции списков, которые следуют после удаляемого списка
	listToDelete, err := listRepository.GetById(IdL)
	if err != nil {
		return err
	}
	err = listRepository.db.Delete(&models.List{}, IdL).Error
	if err != nil {
		return err
	}
	tasks, err := listRepository.GetTasks(IdL)
	for i := range *tasks {
		err = listRepository.db.Where("id_t = ?", (*tasks)[i].IdT).Delete(&models.ImportantTask{}).Error
		if err != nil {
			return err
		}
	}
	return listRepository.db.Model(&models.List{}).
		Where("position > ? AND id_b = ?", listToDelete.Position, listToDelete.IdB).
		UpdateColumn("position", gorm.Expr("position - 1")).Error
}

func (listRepository *ListRepositoryImpl) GetTasks(IdL uint) (*[]models.Task, error) {
	tasks := new([]models.Task)
	result := listRepository.db.Where("id_l = ?", IdL).Order("position").Find(tasks)
	return tasks, result.Error
}

func (listRepository *ListRepositoryImpl) GetById(IdL uint) (*models.List, error) {
	//указатель на структуру, которую вернем
	list := new(models.List)
	result := listRepository.db.Find(list, IdL)
	// если выборка в 0 строк, то такого листа нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrListNotFound
	} else if result.Error != nil {
		// если произошла ошибка при выборке
		return nil, result.Error
	} else {
		// иначе вернем доску
		return list, nil
	}
}

func (listRepository *ListRepositoryImpl) GetBoard(IdL uint) (*models.Board, error) {
	list, err := listRepository.GetById(IdL)
	if err != nil {
		return nil, err
	}
	board := new(models.Board)
	result := listRepository.db.Find(board, list.IdB)
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
