package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func MakeTaskRepository(db *gorm.DB) repositories.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (taskRepository *TaskRepositoryImpl) AppendUser(IdT uint, IdU uint) error {
	user := new(models.User)
	result := taskRepository.db.Find(user, IdU)
	if result.RowsAffected == 0 {
		return customErrors.ErrUserNotFound
	} else if result.Error != nil {
		return result.Error
	}
	err := taskRepository.db.Model(&models.Task{IdT: IdT}).Association("Users").Append(user)
	return err
}

func (taskRepository *TaskRepositoryImpl) Create(task *models.Task, IdL uint, IdB uint) (uint, error) {
	task.IdL = IdL
	task.IdB = IdB
	var currentPosition int64
	err := taskRepository.db.Model(&models.Task{}).Where("id_l = ?", task.IdL).Count(&currentPosition).Error
	if err != nil {
		return 0, err
	}
	task.Position = uint(currentPosition) + 1
	err = taskRepository.db.Create(task).Error
	return task.IdT, err
}

func (taskRepository *TaskRepositoryImpl) GetTasks(IdL uint) (*[]models.Task, error) {
	tasks := new([]models.Task)
	result := taskRepository.db.Where("id_l = ?", IdL).Find(tasks)
	return tasks, result.Error
}

func (taskRepository *TaskRepositoryImpl) Update(task models.Task) error {
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
	if currentData.IsReady != task.IsReady {
		currentData.IsReady = task.IsReady
	}
	if currentData.Deadline != task.Deadline && task.Deadline != "" {
		currentData.Deadline = task.Deadline
	}
	if currentData.IsImportant != task.IsImportant {
		currentData.IsImportant = task.IsImportant
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
	if currentData.Position != task.Position && currentData.IdL != task.IdL && task.IdL != 0 {
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

func (taskRepository *TaskRepositoryImpl) Delete(IdT uint) error {
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

func (taskRepository *TaskRepositoryImpl) GetById(IdT uint) (*models.Task, error) {
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

func (taskRepository *TaskRepositoryImpl) GetCheckLists(IdT uint) (*[]models.CheckList, error) {
	checkLists := new([]models.CheckList)
	err := taskRepository.db.Where("id_t = ?", IdT).Find(checkLists).Error
	return checkLists, err
}

func (taskRepository *TaskRepositoryImpl) IsAccessToTask(IdU uint, IdT uint) (bool, error) {
	task, err := taskRepository.GetById(IdT)
	if err != nil {
		return false, err
	}
	board := new(models.Board)
	err = taskRepository.db.Model(&models.User{IdU: IdU}).Where("id_b = ?", task.IdB).Association("Boards").Find(board)
	if err != nil {
		return false, err
	} else if board == nil {
		return false, nil
	}
	return true, nil
}

func (taskRepository *TaskRepositoryImpl) GetImportantTasks(IdU uint) (*[]models.Task, error) {
	tasks := new([]models.Task)
	err := taskRepository.db.Where("id_u = ? and is_important = true", IdU).Order("date_to_order").Find(tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskRepository *TaskRepositoryImpl) GetTaskUser(IdT uint) (*[]models.User, error) {
	users := new([]models.User)
	err := taskRepository.db.Model(&models.Task{IdT: IdT}).Association("Users").Find(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (taskRepository *TaskRepositoryImpl) DeleteUser(IdT uint, IdU uint) error {
	user := new(models.User)
	result := taskRepository.db.Find(user, IdU)
	if result.RowsAffected == 0 {
		return customErrors.ErrUserNotFound
	} else if result.Error != nil {
		return result.Error
	}
	err := taskRepository.db.Model(&models.Task{IdT: IdT}).Association("Users").Delete(user)
	return err
}
