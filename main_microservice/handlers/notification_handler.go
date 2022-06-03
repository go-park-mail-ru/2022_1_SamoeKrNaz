package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationHandler struct {
	usecase usecases.NotificationUseCase
}

func MakeNotificationHandler(usecase_ usecases.NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{usecase: usecase_}
}

func (notificationHandler *NotificationHandler) GetNotifications(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	//Получаю доски от БД
	var notifications *models.Notifications
	notifications, err := notificationHandler.usecase.GetUsersNotifications(uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	boardsJson, err := notifications.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", boardsJson)
}

func (notificationHandler *NotificationHandler) ReadNotifications(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	//Получаю доски от БД
	err := notificationHandler.usecase.ReadNotifications(uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}
	var isOkay models.Is_okayIn
	isOkay.Is_okayInfo = true
	isOkayJson, err := isOkay.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", isOkayJson)
}

func (notificationHandler *NotificationHandler) DeleteNotifications(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	//Получаю доски от БД
	err := notificationHandler.usecase.DeleteNotifications(uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}
	var isOkay models.Is_okayIn
	isOkay.Is_okayInfo = true
	isOkayJson, err := isOkay.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", isOkayJson)
}
