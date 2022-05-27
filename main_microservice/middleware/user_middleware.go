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
	redisRep repositories.SessionRepository
	boardRep repositories.BoardRepository

	ws wsplanexa.WebSocketPool
}

func CreateMiddleware(redisRep repositories.SessionRepository, boardRep repositories.BoardRepository) *Middleware {
	return &Middleware{redisRep: redisRep, boardRep: boardRep}
}

func (mw *Middleware) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД
	userId, err := mw.redisRep.GetSession(token)
	if err != nil {
		return
	}
	c.Set("Auth", userId)
}

func (mw *Middleware) SendToWebSocket(c *gin.Context) {
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
	} else if event.EventType == "UpdateTask" {
		currentIdT, check := c.Get("IdT")
		if !check {
			_ = c.Error(customErrors.ErrBadInputData)
			return
		}
		event.IdT = uint(currentIdT.(uint64))
	} else {
		return
	}
	boardsUsers, err := mw.boardRep.GetBoardUser(event.IdB)
	if err != nil {
		_ = c.Error(err)
		return
	}
	for _, user := range boardsUsers {
		if user.IdU == userId {
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
