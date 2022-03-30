package repositories

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func (taskRepository *TaskRepository) MakeRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (taskRepository *TaskRepository) Create(task *models.Task, IdL uint, IdB uint) error {
	task.IdL = IdL
	task.IdB = IdB
	var currentPosition int64
	result := taskRepository.db.Model(&models.Task{}).Where("id_t = ?", task.IdT).Count(&currentPosition)
	if result.Error != nil {
		return result.Error
	}
	task.Position = uint(currentPosition) + 1
	return taskRepository.db.Create(task).Error
}

func (taskRepository *TaskRepository) Update(task *models.Task) error {
	currentData, err := taskRepository.GetById(task.IdT)
	if err != nil {
		return err
	}
	if currentData.Title != task.Title {
		currentData.Title = task.Title
	}
	if currentData.Description != task.Description {
		currentData.Description = task.Description
	}
	if currentData.Position != task.Position {
		currentData.Position = task.Position
	}
	return taskRepository.db.Save(currentData).Error
}

func (taskRepository *TaskRepository) Delete(IdT uint) error {
	return taskRepository.db.Delete(&models.Task{}, IdT).Error
}

func (taskRepository *TaskRepository) GetById(IdT uint) (*models.Task, error) {
	// указатель на структуру, которую вернем
	task := new(models.Task)
	result := taskRepository.db.Find(task, IdT)
	// если выборка в 0 строк, то такой доски нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrBoardNotFound
	} else if result.Error != nil {
		// если произошла ошибка при выборке
		return nil, result.Error
	} else {
		// иначе вернем доску
		return task, nil
	}
}
