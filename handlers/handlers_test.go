package handlers

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/routes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
)

var (
	router = gin.New()
	lock   = sync.RWMutex{}
)

func TestGetBoards(t *testing.T) {
	t.Parallel()

	request, _ := http.NewRequest("GET", routes.HomeRoute, nil)
	cookie := &http.Cookie{
		Name:  "1",
		Value: "session1",
	}
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Error("status is not ok")
	}

	var returnedBoards []models.Board
	err := json.Unmarshal(writer.Body.Bytes(), &returnedBoards)
	if err != nil {
		t.Error(err)
	}

	lock.RLock()
	expectedBoards := models.BoardList
	isEqual := true

	if len(expectedBoards) != len(returnedBoards) {
		isEqual = false
	}

	if !reflect.DeepEqual(returnedBoards, expectedBoards) {
		isEqual = false
	}
	lock.RUnlock()

	require.True(t, isEqual)
}
