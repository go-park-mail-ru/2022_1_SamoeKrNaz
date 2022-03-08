package handlers

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/routes"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var (
	router = gin.New()
)

func TestMain(m *testing.M) {
	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.LoginRoute, Login)
		mainRoutes.POST(routes.RegisterRoute, Register)
		mainRoutes.GET("", GetBoards)
	}
	os.Exit(m.Run())
}

func TestGetBoardsSuccess(t *testing.T) {
	t.Parallel()

	request, _ := http.NewRequest("GET", routes.HomeRoute, nil)
	cookie := &http.Cookie{
		Name:  "token",
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

	expectedBoards := models.BoardList
	isEqual := true

	if !reflect.DeepEqual(returnedBoards, expectedBoards) {
		isEqual = false
	}

	require.True(t, isEqual)
}

func TestGetBoardsFalse(t *testing.T) {
	t.Parallel()

	request, _ := http.NewRequest("GET", routes.HomeRoute, nil)
	cookie := &http.Cookie{
		Name:  "token",
		Value: "sessionFalse",
	}
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	require.Equal(t, writer.Code, http.StatusUnauthorized)
}

func TestLoginSuccess(t *testing.T) {
	t.Parallel()

	newUser := models.User{Username: "paperThing11", Password: "gedab1gawf"}
	jsonNewUser, _ := json.Marshal(newUser)
	body := bytes.NewReader(jsonNewUser)

	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.LoginRoute, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	require.Equal(t, http.StatusOK, writer.Code)
}

func TestLoginFail(t *testing.T) {
	t.Parallel()

	newUser := models.User{Username: "user123", Password: "pass123"}
	jsonNewUser, _ := json.Marshal(newUser)
	body := bytes.NewReader(jsonNewUser)

	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.LoginRoute, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	require.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRegisterSuccess(t *testing.T) {
	t.Parallel()

	newUser := models.User{Username: "zxc_god", Password: "kaneki_ken"}
	jsonNewUser, _ := json.Marshal(newUser)
	body := bytes.NewReader(jsonNewUser)

	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.RegisterRoute, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	require.Equal(t, http.StatusCreated, writer.Code)

	isEqual := false

	for _, user := range models.UserList {
		if user.Username == newUser.Username && user.Password == newUser.Password {
			isEqual = true
		}
	}

	require.True(t, isEqual)
}

func TestRegisterFail(t *testing.T) {
	t.Parallel()

	newUser := models.User{Username: "palantina14", Password: "bdazglweq21"}
	jsonNewUser, _ := json.Marshal(newUser)
	body := bytes.NewReader(jsonNewUser)

	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.RegisterRoute, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	require.Equal(t, http.StatusConflict, writer.Code)
}

func TestRegisterBadPassword(t *testing.T) {
	t.Parallel()

	newUser := models.User{Username: "cucumber_two_two", Password: "я люблю Россию"}
	jsonNewUser, _ := json.Marshal(newUser)
	body := bytes.NewReader(jsonNewUser)

	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.RegisterRoute, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	require.Equal(t, http.StatusBadRequest, writer.Code)
}

func TestLogoutSuccess(t *testing.T) {
	t.Parallel()

	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.LogoutRoute, nil)
	cookie := &http.Cookie{
		Name:  "token",
		Value: "session1",
	}
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	require.Equal(t, http.StatusOK, writer.Code)
}
