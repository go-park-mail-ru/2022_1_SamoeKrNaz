package middleware

import (
	"PLANEXA_backend/main_microservice/repositories"
	wsplanexa "PLANEXA_backend/main_microservice/websocket"
	"PLANEXA_backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware struct {
	sessionRepo repositories.SessionRepository
	boardRepo   repositories.BoardRepository

	ws wsplanexa.WebSocketPool
}

func CreateMiddleware(sessionRepo repositories.SessionRepository, boardRepo repositories.BoardRepository, ws wsplanexa.WebSocketPool) *Middleware {
	return &Middleware{sessionRepo: sessionRepo, boardRepo: boardRepo, ws: ws}
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

func (mw *Middleware) SendNotification(c *gin.Context) {
	c.Next()
	status := c.Writer.Status()
	if status != http.StatusOK && status != http.StatusCreated {
		fmt.Println("error in httpstatus")
		return
	}

	currentEvent, check := c.Get("eventType")
	if !check {
		fmt.Println("error in get eventType")
		return
	}

	event := &models.Event{
		EventType: currentEvent.(string),
	}

	if event.EventType == "AppendUserToBoard" {
		currentIdB, check := c.Get("IdB")
		if !check {
			fmt.Println("error in get idb")
			return
		}
		event.IdB = currentIdB.(uint)
		boardsUsers, err := mw.boardRepo.GetBoardUser(event.IdB)
		if err != nil {
			fmt.Println("error in getboarduser")
			return
		}
		for _, user := range boardsUsers {
			eventJson, err := event.MarshalJSON()
			if err != nil {
				fmt.Println("error in marshalljson")
				return
			}
			mw.ws.Send(user.IdU, eventJson)
		}
	}
}

func (mw *Middleware) SendToWebSocket(c *gin.Context) {
	c.Next()
	status := c.Writer.Status()
	if status != http.StatusOK && status != http.StatusCreated {
		fmt.Println("error in httpstatus")
		return
	}

	currentEvent, check := c.Get("eventType")
	if !check {
		fmt.Println("error in get eventType")
		return
	}
	event := &models.Event{
		EventType: currentEvent.(string),
	}

	userId, check := c.Get("Auth")
	if !check {
		fmt.Println("error in get auth")
		return
	}

	currentIdB, check := c.Get("IdB")
	if !check {
		fmt.Println("error in get idb")
		return
	}
	event.IdB = currentIdB.(uint)

	if event.EventType == "UpdateTask" || event.EventType == "DeleteTask" {
		currentIdT, check := c.Get("IdT")
		if !check {
			fmt.Println("error in get idt")
			return
		}
		event.IdT = currentIdT.(uint)
	}

	boardsUsers, err := mw.boardRepo.GetBoardUser(event.IdB)

	if err != nil {
		fmt.Println("error in getboarduser")
		return
	}
	for _, user := range boardsUsers {
		if user.IdU == uint(userId.(uint64)) {
			continue
		}
		eventJson, err := event.MarshalJSON()
		if err != nil {
			fmt.Println("error in marshalljson")
			return
		}
		mw.ws.Send(user.IdU, eventJson)
	}
}
