package impl

import (
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	rtime "github.com/ivahaev/russian-time"
	"strconv"
	"time"
)

type NotificationUseCaseImpl struct {
	repNotification repositories.NotificationRepository
	repBoard        repositories.BoardRepository
	repTask         repositories.TaskRepository
	repUser         repositories.UserRepository
}

func MakeNotificationUseCase(repNotification_ repositories.NotificationRepository,
	repBoard_ repositories.BoardRepository, repTask_ repositories.TaskRepository, repUser_ repositories.UserRepository) usecases.NotificationUseCase {
	return &NotificationUseCaseImpl{repNotification: repNotification_,
		repBoard: repBoard_, repTask: repTask_, repUser: repUser_}
}

func (notificationUsecase NotificationUseCaseImpl) Create(notification *models.Notification) error {
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return err
	}
	notification.Date = strconv.Itoa(time.Now().In(moscow).Day()) + " " + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + time.Now().In(moscow).Format("15:04")
	if notification.IdB != 0 {
		currentBoard, err := notificationUsecase.repBoard.GetById(notification.IdB)
		if err != nil {
			return err
		}
		notification.Board = *currentBoard
	}
	if notification.IdT != 0 {
		currentTask, err := notificationUsecase.repTask.GetById(notification.IdT)
		if err != nil {
			return err
		}
		notification.Task = *currentTask
		if notification.IdB == 0 {
			currentBoard, err := notificationUsecase.repBoard.GetById(currentTask.IdB)
			if err != nil {
				return nil
			}
			notification.Board = *currentBoard
			notification.IdB = (*currentBoard).IdB
		}
	}
	notification.DateToOrder = time.Now()
	if notification.IdWh != 0 {
		currentUser, err := notificationUsecase.repUser.GetUserById(notification.IdWh)
		if err != nil {
			return err
		}
		notification.UserWho = *currentUser
	}
	err = notificationUsecase.repNotification.Create(notification)
	if err != nil {
		return err
	}
	return nil
}

func (notificationUsecase NotificationUseCaseImpl) CreateBoardNotification(notification *models.Notification) error {
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return err
	}
	notification.Date = strconv.Itoa(time.Now().In(moscow).Day()) + " " + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + time.Now().In(moscow).Format("15:04")
	currentBoard, err := notificationUsecase.repBoard.GetById(notification.IdB)
	if err != nil {
		return err
	}
	notification.DateToOrder = time.Now()
	notification.Board = *currentBoard
	userFromBoard, err := notificationUsecase.repBoard.GetBoardUser(notification.IdB)
	if err != nil {
		return err
	}
	currentUser, err := notificationUsecase.repUser.GetUserById(notification.IdWh)
	if err != nil {
		return err
	}
	notification.UserWho = *currentUser
	//каждому нужно создать уведомление, что был добавлен пользователь
	for _, user := range userFromBoard {
		if user.IdU == notification.IdWh {
			continue
		}
		notification.IdU = user.IdU
		err := notificationUsecase.repNotification.Create(notification)
		if err != nil {
			return err
		}
	}
	return nil
}

func (notificationUsecase NotificationUseCaseImpl) GetUsersNotifications(IdU uint) (*models.Notifications, error) {
	notifications, err := notificationUsecase.repNotification.GetUsersNotifications(IdU)
	if err != nil {
		return nil, err
	}
	for i := range *notifications {
		if (*notifications)[i].IdB != 0 {
			currentBoard, err := notificationUsecase.repBoard.GetById((*notifications)[i].IdB)
			if err != nil {
				return nil, err
			}
			(*notifications)[i].Board = *currentBoard
		}
		if (*notifications)[i].IdT != 0 {
			currentTask, err := notificationUsecase.repTask.GetById((*notifications)[i].IdT)
			if err != nil {
				return nil, err
			}
			(*notifications)[i].Task = *currentTask
		}
		if (*notifications)[i].IdWh != 0 {
			currentUser, err := notificationUsecase.repUser.GetUserById((*notifications)[i].IdWh)
			if err != nil {
				return nil, err
			}
			(*notifications)[i].UserWho = *currentUser
		}
	}
	return (*models.Notifications)(notifications), nil
}

func (notificationUsecase NotificationUseCaseImpl) ReadNotifications(IdU uint) error {
	err := notificationUsecase.repNotification.ReadNotifications(IdU)
	if err != nil {
		return err
	}
	return nil
}

func (notificationUsecase NotificationUseCaseImpl) DeleteNotifications(IdU uint) error {
	err := notificationUsecase.repNotification.DeleteNotifications(IdU)
	if err != nil {
		return err
	}
	return nil
}
