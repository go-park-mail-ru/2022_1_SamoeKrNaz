package repositories

import "PLANEXA_backend/models"

type TaskRepository interface {
	AppendUser(IdT uint, IdU uint) error
	Create(task *models.Task, IdL uint, IdB uint) (uint, error)
	GetTasks(IdL uint) (*[]models.Task, error)
	Update(task models.Task) error
	Delete(IdT uint) error
	GetById(IdT uint) (*models.Task, error)
	GetCheckLists(IdT uint) (*[]models.CheckList, error)
	IsAccessToTask(IdU uint, IdT uint) (bool, error)
	GetImportantTasks(IdU uint) (*[]models.Task, error)
	GetTaskUser(IdT uint) (*[]models.User, error)
	DeleteUser(IdT uint, IdU uint) error
	GetAttachments(IdT uint) (*[]models.Attachment, error)
	GetByLink(link string) (*models.Task, error)
	IsAppendedToTask(IdU uint, IdT uint) (bool, error)
}
