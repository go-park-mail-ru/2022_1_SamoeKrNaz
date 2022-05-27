package wsplanexa_impl

import (
	customErrors "PLANEXA_backend/errors"
	wsplanexa "PLANEXA_backend/main_microservice/websocket"
	"PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Pool struct {
	socketPool map[uint][]*websocket.Conn
	mutex      *sync.Mutex
	upgrader   websocket.Upgrader
}

func CreatePool() wsplanexa.WebSocketPool {
	pool := &Pool{
		socketPool: make(map[uint][]*websocket.Conn),
		mutex:      &sync.Mutex{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	return pool
}

func (pool *Pool) Start(c *gin.Context) {

	ws, err := pool.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		_ = c.Error(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			_ = c.Error(err)
			return
		}
	}(ws)
	userId, check := c.Get("IdU")
	if !check {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	err = pool.Add(uint(userId.(uint64)), ws)
	if err != nil {
		_ = c.Error(err)
		return
	}
	defer func() {
		err := pool.Delete(uint(userId.(uint64)), ws)
		if err != nil {
			_ = c.Error(err)
			return
		}
	}()
	var errorRead error = nil
	for errorRead == nil {
		_, _, errorRead = ws.ReadMessage()
		if errorRead != nil {
			_ = c.Error(err)
			return
		}
	}
	var isAppended models.Appended
	isAppended.AppendedInfo = true
	isAppendedJson, err := isAppended.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", isAppendedJson)
}

func (pool *Pool) Add(IdU uint, ws *websocket.Conn) error {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	for _, item := range pool.socketPool[IdU] {
		if item == ws {
			return customErrors.ErrWebSocketExist
		}
	}
	pool.socketPool[IdU] = append(pool.socketPool[IdU], ws)
	return nil
}

func (pool *Pool) Delete(IdU uint, ws *websocket.Conn) error {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	for i, item := range pool.socketPool[IdU] {
		if item == ws {
			pool.socketPool[IdU][i] = pool.socketPool[IdU][len(pool.socketPool[IdU])-1]
			pool.socketPool[IdU] = pool.socketPool[IdU][:len(pool.socketPool[IdU])-1]
			return nil
		}
	}
	return customErrors.ErrWebSocketNotFound
}

func (pool *Pool) Send(IdU uint, data []byte) error {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	for _, item := range pool.socketPool[IdU] {
		err := item.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return err
		}
	}
	return nil
}
