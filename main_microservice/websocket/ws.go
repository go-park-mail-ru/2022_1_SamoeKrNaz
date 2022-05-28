package wsplanexa

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketPool interface {
	Start(c *gin.Context)
	Add(IdU uint, ws *websocket.Conn) error
	Delete(IdU uint, ws *websocket.Conn) error
	Send(IdU uint, data []byte)
}
