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
	result := listRepository.db.Model(&models.List{}).Where("id_b = ?", list.IdB).Count(&currentPosition)
	if result.Error != nil {
		return result.Error
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
		currentData.Position = list.Position
	}
	//сохраняем новую структуру
	return listRepository.db.Save(currentData).Error
}
func (listRepository *ListRepository) Delete(IdL uint) error {
	return listRepository.db.Delete(&models.List{}, IdL).Error
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
