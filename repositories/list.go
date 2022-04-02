package repositories

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
)

type ListRepository struct {
	db *gorm.DB
}

func (listRepository *ListRepository) MakeRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (listRepository *ListRepository) Create(list *models.List, IdB uint) error {
	list.IdB = IdB
	var currentPosition int64
	err := listRepository.db.Model(&models.List{}).Where("id_b = ?", list.IdB).Count(&currentPosition).Error
	if err != nil {
		return err
	}
	list.Position = uint(currentPosition) + 1
	return listRepository.db.Create(list).Error
}

func (listRepository *ListRepository) Update(list *models.List) error {
	currentData, err := listRepository.GetById(list.IdL)
	if err != nil {
		return err
	}
	if currentData.Title != list.Title {
		currentData.Title = list.Title
	}
	if currentData.Position != list.Position {
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
func (listRepository *ListRepository) Delete(IdL uint) error {
	// при удалении необходимо изменить позиции списков, которые следуют после удаляемого списка
	listToDelete, err := listRepository.GetById(IdL)
	if err != nil {
		return err
	}
	err = listRepository.db.Delete(&models.List{}, IdL).Error
	if err != nil {
		return err
	}
	return listRepository.db.Model(&models.Task{}).
		Where("position > ? AND id_b = ?", listToDelete.Position, listToDelete.IdB).
		UpdateColumn("position", gorm.Expr("position - 1")).Error
}

func (listRepository *ListRepository) GetTasks(IdL uint) (*[]models.Task, error) {
	tasks := new([]models.Task)
	result := listRepository.db.Where("id_l = ?", IdL).Find(tasks)
	return tasks, result.Error
}

func (listRepository *ListRepository) GetById(IdL uint) (*models.List, error) {
	// указатель на структуру, которую вернем
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

func (listRepository *ListRepository) ChangePosition(currentPosition, newPosition uint) {

}