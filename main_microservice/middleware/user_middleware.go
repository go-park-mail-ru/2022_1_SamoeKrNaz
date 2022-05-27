package middleware

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	wsplanexa "PLANEXA_backend/main_microservice/websocket"
	"PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware struct {
	sessionRepo repositories.SessionRepository
	boardRepo   repositories.BoardRepository

	ws wsplanexa.WebSocketPool
}

func CreateMiddleware(sessionRepo repositories.SessionRepository, boardRepo repositories.BoardRepository) *Middleware {
	return &Middleware{sessionRepo: sessionRepo, boardRepo: boardRepo}
}

func (mw *Middleware) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД
	userId, err := mw.sessionRepo.GetSession(token)
	if err != nil {
		return
	}
	c.Set("Auth", userId)
}

func (mw *Middleware) SendToWebSocket(c *gin.Context) {
	c.Next()
	status := c.Writer.Status()
	if status != http.StatusOK {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	currentEvent, check := c.Get("eventType")
	if !check {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	event := &models.Event{
		EventType: currentEvent.(string),
	}
	userId, check := c.Get("IdU")
	if !check {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	if event.EventType == "UpdateBoard" {
		currentIdB, check := c.Get("IdB")
		if !check {
			_ = c.Error(customErrors.ErrBadInputData)
			return
		}
		event.IdB = uint(currentIdB.(uint64))
	} else if event.EventType == "UpdateTask" || event.EventType == "DeleteTask" {
		currentIdT, check := c.Get("IdT")
		if !check {
			_ = c.Error(customErrors.ErrBadInputData)
			return
		}
		event.IdT = uint(currentIdT.(uint64))
	} else {
		return
	}
	boardsUsers, err := mw.boardRepo.GetBoardUser(event.IdB)
	if err != nil {
		_ = c.Error(err)
		return
	}
	for _, user := range boardsUsers {
		if user.IdU == uint(userId.(uint64)) {
			continue
		}
		eventJson, err := event.MarshalJSON()
		if err != nil {
			_ = c.Error(err)
			return
		}
		err = mw.ws.Send(user.IdU, eventJson)
		if err != nil {
			_ = c.Error(err)
			return
		}
	}
}
