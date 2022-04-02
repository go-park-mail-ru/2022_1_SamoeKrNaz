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
	result := taskRepository.db.Model(&models.Task{}).Where("id_l = ?", task.IdL).Count(&currentPosition).Error
	if result != nil {
		return result
	}
	task.Position = uint(currentPosition) + 1
	return taskRepository.db.Create(task).Error
}

func (taskRepository *TaskRepository) Update(task *models.Task) error {
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
	if currentData.Position != task.Position {
		repository := ListRepository{}
		listRepo := repository.MakeRepository(taskRepository.db)
		// получим все списки из текущей доски
		taskInList, err := listRepo.GetTasks(currentData.IdL)
		if err != nil {
			return err
		}
		// если список переместили вниз
		if currentData.Position > task.Position {
			for i := task.Position - 1; i < currentData.Position-1; i++ {
				(*taskInList)[i].Position += 1
				(*taskInList)[i].IdT += 1
				err = taskRepository.db.Save((*taskInList)[i]).Error
				if err != nil {
					return err
				}
			}
			currentData.Position = task.Position
			currentData.IdT = task.Position
		} else { // если список переместили вверх
			for i := currentData.Position; i < task.Position; i++ {
				(*taskInList)[i].Position -= 1
				(*taskInList)[i].IdT -= 1
				err = taskRepository.db.Save((*taskInList)[i]).Error
				if err != nil {
					return err
				}
			}
			currentData.Position = task.Position
			currentData.IdT = task.Position
		}
	}
	return taskRepository.db.Save(currentData).Error
}

func (taskRepository *TaskRepository) Delete(IdT uint) error {
	// при удалении необходимо изменить позиции тасков, которые следуют после удаляемой задачи
	taskToDelete, err := taskRepository.GetById(IdT)
	if err != nil {
		return err
	}
	repository := ListRepository{}
	taskRepo := repository.MakeRepository(taskRepository.db)
	// получим все таски из текущего списка
	listsInBoards, err := taskRepo.GetTasks(taskToDelete.IdB)
	if err != nil {
		return err
	}
	err = taskRepository.db.Delete(&models.Task{}, IdT).Error
	if err != nil {
		return err
	}
	for i := int(taskToDelete.Position); i < len(*listsInBoards); i++ {
		// сдвинем позицию на одну
		(*listsInBoards)[i].Position -= 1
		// и удалим
		err = taskRepository.db.Save((*listsInBoards)[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (taskRepository *TaskRepository) GetById(IdT uint) (*models.Task, error) {
	// указатель на структуру, которую вернем
	task := new(models.Task)
	result := taskRepository.db.Find(task, IdT)
	// если выборка в 0 строк, то такой таски нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrBoardNotFound
	} else if result.Error != nil {
		// если произошла ошибка при выборке
		return nil, result.Error
	} else {
		// иначе вернем таску
		return task, nil
	}
}
