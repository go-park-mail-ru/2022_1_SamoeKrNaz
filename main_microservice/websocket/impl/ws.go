package wsplanexa_impl

import (
	customErrors "PLANEXA_backend/errors"
	wsplanexa "PLANEXA_backend/main_microservice/websocket"
	"PLANEXA_backend/models"
	"fmt"
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
	defer ws.Close()
	userId, check := c.Get("Auth")
	if !check {
		fmt.Println("error in get auth")
		return
	}
	err = pool.Add(uint(userId.(uint64)), ws)
	if err != nil {
		fmt.Println("error in add")
		return
	}
	defer func() {
		err := pool.Delete(uint(userId.(uint64)), ws)
		if err != nil {
			fmt.Println("error in delete")
			return
		}
	}()
	var errorRead error = nil
	for errorRead == nil {
		_, _, errorRead = ws.ReadMessage()
		if errorRead != nil {
			fmt.Println("error in errorRead")
			return
		}
	}
	var isDeleted models.Deleted
	isDeleted.DeletedInfo = true
	isDeletedJson, err := isDeleted.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
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

func (pool *Pool) Send(IdU uint, data []byte) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	for _, item := range pool.socketPool[IdU] {
		err := item.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
