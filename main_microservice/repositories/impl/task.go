package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	var currentPosition int64
	err := taskRepository.db.Model(&models.Task{}).Where("id_l = ?", task.IdL).Count(&currentPosition).Error
	if err != nil {
		return 0, err
	}
	task.Position = uint(currentPosition) + 1
	task.IconPattern = uint(rand.Intn(5) + 1)
	err = taskRepository.db.Create(task).Error
	return task.IdT, err
}

func (taskRepository *TaskRepositoryImpl) GetTasks(IdL uint, IdU uint) (*[]models.Task, error) {
	tasks := new([]models.Task)
	result := taskRepository.db.Where("id_l = ?", IdL).Find(tasks)
	for i := range *tasks {
		(*tasks)[i].IsImportant = taskRepository.GetImportance((*tasks)[i].IdT, IdU)
	}
	return tasks, result.Error
}

func (taskRepository *TaskRepositoryImpl) Update(task models.Task, IdU uint) error {
	currentData, err := taskRepository.GetById(task.IdT)
	if err != nil {
		return err
	}
	if currentData.Title != task.Title && task.Title != "" {
		currentData.Title = task.Title
	}
	if currentData.Description != task.Description && task.Description != "" {
		currentData.Description = task.Description
	}
	if currentData.IsReady != task.IsReady {
		currentData.IsReady = task.IsReady
	}
	if currentData.Deadline != task.Deadline && task.Deadline != "" {
		currentData.Deadline = task.Deadline
	}
	if task.IsImportant != "" {
		if task.IsImportant == "false" {
			err := taskRepository.db.Delete(&models.ImportantTask{IdB: currentData.IdB, IdT: currentData.IdT, IdU: IdU}).Error
			if err != nil {
				return err
			}
		} else {
			err := taskRepository.db.Create(&models.ImportantTask{IdB: currentData.IdB, IdT: currentData.IdT, IdU: IdU}).Error
			if err != nil {
				return err
			}
		}
	}
	if currentData.Position != task.Position && currentData.IdL == task.IdL {
		// если список переместили вниз
		if currentData.Position > task.Position {
			// допустим, что был список 1 2 3 4
			// решили, что четвертый список будет после первого
			// 1 4 2 3
			// значит, нужно все индексы после текущей позиции увеличить на 1
			err := taskRepository.db.Model(&models.Task{}).
				Where("id_l = ? AND position BETWEEN ? AND ?", currentData.IdL, task.Position, currentData.Position-1).
				UpdateColumn("position", gorm.Expr("position + 1")).Error
			if err != nil {
				return err
			}
			currentData.Position = task.Position
		} else { // если список переместили вверх
			// допустим, что был список 1 2 3 4
			// решили, что второй список будет после четвертого
			// 1 3 4 2
			// значит, нужно все индексы  с предыдущей позиции уменьшить на 1
			err := taskRepository.db.Model(&models.Task{}).
				Where("id_l = ? AND position BETWEEN ? AND ?", currentData.IdL, currentData.Position+1, task.Position).
				UpdateColumn("position", gorm.Expr("position - 1")).Error
			if err != nil {
				return err
			}
			currentData.Position = task.Position
		}
	}
	if currentData.IdL != task.IdL && task.IdL != 0 {
		fmt.Println("второе условие")
		// если мы переместили таску из одного списка в другой и поменяли список
		// то нужно в старом списке поменять позиции после текущей таски
		err := taskRepository.db.Model(&models.Task{}).
			Where("position > ? AND id_l = ?", currentData.Position, currentData.IdL).
			UpdateColumn("position", gorm.Expr("position - 1")).Error
		if err != nil {
			return err
		}
		// в новом списке поменять позицию, то не будем учитывать позицию в прошлом листе: новый лист - новые позиции
		// будем менять в том случае, если в новом листе после этой позиции есть таски
		// иначе просто сохраняем с новым индексом
		err = taskRepository.db.Model(&models.Task{}).
			Where("id_l = ? AND position >= ?", task.IdL, task.Position).
			UpdateColumn("position", gorm.Expr("position + 1")).Error
		if err != nil {
			return err
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
	err = taskRepository.db.Where("id_t = ?", IdT).Delete(&models.Notification{}).Error
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

func (taskRepository *TaskRepositoryImpl) GetImportance(IdT uint, IdU uint) string {
	importantTask := new(models.ImportantTask)
	importantResult := taskRepository.db.Where("id_u = ? and id_t = ?", IdU, IdT).Find(importantTask)
	// если выборка в 0 строк, то такой таски нет
	if importantResult.RowsAffected == 0 {
		return "false"
	} else {
		return "true"
	}
}

func (taskRepository *TaskRepositoryImpl) GetCheckLists(IdT uint) (*[]models.CheckList, error) {
	checkLists := new([]models.CheckList)
	err := taskRepository.db.Where("id_t = ?", IdT).Order("id_cl").Find(checkLists).Error
	return checkLists, err
}

func (taskRepository *TaskRepositoryImpl) IsAccessToTask(IdU uint, IdT uint) (bool, error) {
	task, err := taskRepository.GetById(IdT)
	if err != nil {
		return false, err
	}
	user := new(models.User)
	err = taskRepository.db.Model(&models.Board{IdB: task.IdB}).Where("id_u = ?", IdU).Association("Users").Find(user)
	if err != nil {
		return false, err
	} else if user.IdU == 0 {
		return false, nil
	}
	return true, nil
}

func (taskRepository *TaskRepositoryImpl) IsAppendedToTask(IdU uint, IdT uint) (bool, error) {
	user := new(models.User)
	err := taskRepository.db.Model(&models.Task{IdT: IdT}).Where("id_u = ?", IdU).Association("Users").Find(user)
	if err != nil {
		return false, err
	} else if user.IdU == 0 {
		return false, nil
	}
	return true, nil
}

func (taskRepository *TaskRepositoryImpl) GetImportantTasks(IdU uint) (*[]models.Task, error) {
	importantTasks := new([]models.ImportantTask)
	err := taskRepository.db.Where("id_u = ?", IdU).Order("date_to_order").Find(importantTasks).Error
	if err != nil {
		return nil, err
	}
	tasks := new([]models.Task)
	for i := range *importantTasks {
		currentImportant := new([]models.Task)
		err = taskRepository.db.Where("id_t = ?", (*importantTasks)[i].IdT).Find(currentImportant).Error
		if err != nil {
			return nil, err
		}
		*tasks = append(*tasks, *currentImportant...)
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

func (taskRepository *TaskRepositoryImpl) GetAttachments(IdT uint) (*[]models.Attachment, error) {
	attachments := new([]models.Attachment)
	err := taskRepository.db.Where("id_t = ?", IdT).Order("id_a").Find(attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (taskRepository *TaskRepositoryImpl) GetByLink(link string) (*models.Task, error) {
	// указатель на структуру, которую вернем
	task := new(models.Task)
	err := taskRepository.db.Where("link = ?", link).Find(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}
