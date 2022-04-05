package repositories

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func MakeTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (taskRepository *TaskRepository) Create(task models.Task, IdL uint, IdB uint) error {
	task.IdL = IdL
	task.IdB = IdB
	var currentPosition int64
	err := taskRepository.db.Model(&models.Task{}).Where("id_l = ?", task.IdL).Count(&currentPosition).Error
	if err != nil {
		return err
	}
	task.Position = uint(currentPosition) + 1
	return taskRepository.db.Create(task).Error
}

func (taskRepository *TaskRepository) GetTasks(IdL uint) (*[]models.Task, error) {
	tasks := new([]models.Task)
	result := taskRepository.db.Where("id_b = ?", IdL).Find(tasks)
	return tasks, result.Error
}

func (taskRepository *TaskRepository) Update(task models.Task) error {
	currentData, err := taskRepository.GetById(task.IdT)
	if err != nil {
		return err
	}
	if currentData.Title != task.Title && task.Title != "" {
		currentData.Title = task.Title
	}
	if currentData.Description != task.Description && task.Title != "" {
		currentData.Description = task.Description
	}
	if currentData.Position != task.Position && currentData.IdL == task.IdL {
		// если список переместили вниз
		if currentData.Position > task.Position {
			err := taskRepository.db.Model(&models.Task{}).
				Where("id_b = ? AND position BETWEEN ? AND ?", currentData.IdL, task.Position, currentData.Position-1).
				UpdateColumn("position", gorm.Expr("position + 1")).Error
			if err != nil {
				return err
			}
		} else { // если список переместили вверх
			err := taskRepository.db.Model(&models.Task{}).
				Where("id_b = ? AND position BETWEEN ? AND ?", currentData.IdL, currentData.Position+1, task.Position).
				UpdateColumn("position", gorm.Expr("position - 1")).Error
			if err != nil {
				return err
			}
		}
		currentData.Position = task.Position
	}
	if currentData.Position != task.Position && currentData.IdL != task.IdL {
		// если мы переместили таску из одного списка в другой и поменяли список
		// то нужно в старом списке поменять позиции после текущей таски
		err := taskRepository.db.Model(&models.Task{}).
			Where("position > ? AND id_l = ?", currentData.Position, currentData.IdL).
			UpdateColumn("position", gorm.Expr("position - 1")).Error
		if err != nil {
			return err
		}
		// в новом списке поменять позицию
		if currentData.Position > task.Position {
			err := taskRepository.db.Model(&models.Task{}).
				Where("id_b = ? AND position BETWEEN ? AND ?", task.IdL, task.Position, currentData.Position-1).
				UpdateColumn("position", gorm.Expr("position + 1")).Error
			if err != nil {
				return err
			}
		} else { // если список переместили вверх
			err := taskRepository.db.Model(&models.Task{}).
				Where("id_b = ? AND position BETWEEN ? AND ?", task.IdL, currentData.Position+1, task.Position).
				UpdateColumn("position", gorm.Expr("position - 1")).Error
			if err != nil {
				return err
			}
		}
		currentData.Position = task.Position
		currentData.IdL = task.IdL
	}
	return taskRepository.db.Save(currentData).Error
}

func (taskRepository *TaskRepository) Delete(IdT uint) error {
	// при удалении необходимо изменить позиции тасков, которые следуют после удаляемой задачи
	taskToDelete, err := taskRepository.GetById(IdT)
	if err != nil {
		return err
	}
	err = taskRepository.db.Delete(&models.Task{}, IdT).Error
	if err != nil {
		return err
	}
	return taskRepository.db.Model(&models.Task{}).
		Where("position > ? AND id_l = ?", taskToDelete.Position, taskToDelete.IdL).
		UpdateColumn("position", gorm.Expr("position - 1")).Error
}

func (taskRepository *TaskRepository) GetById(IdT uint) (*models.Task, error) {
	// указатель на структуру, которую вернем
	task := new(models.Task)
	result := taskRepository.db.Find(task, IdT)
	// если выборка в 0 строк, то такой таски нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrTaskNotFound
	} else if result.Error != nil {
		// если произошла ошибка при выборке
		return nil, result.Error
	} else {
		// иначе вернем таску
		return task, nil
	}
}
