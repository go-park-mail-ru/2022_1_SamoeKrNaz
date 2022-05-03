package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/middleware"
	"PLANEXA_backend/models"
	mock_repositories "PLANEXA_backend/repositories/mocks"
	"PLANEXA_backend/routes"
	"PLANEXA_backend/usecases/mocks"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasks(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	taskUseCase := mock_usecases.NewMockTaskUseCase(controller)
	taskHandler := MakeTaskHandler(taskUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.ListRoute+"/:id"+routes.TaskRoute, authMiddleware.CheckAuth, taskHandler.GetTasks)
	}
	tasks := []models.Task{{Title: "title1", Description: "desc1"}, {Title: "title2", Description: "desc2"}}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	taskUseCase.EXPECT().GetTasks(uint(11), uint(22)).Return(tasks, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.ListRoute+"/11"+routes.TaskRoute, nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newTasks []models.Task
	_ = json.Unmarshal(writer.Body.Bytes(), &newTasks)
	assert.Equal(t, tasks, newTasks)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.ListRoute+"/11"+routes.TaskRoute, nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetTask(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	taskUseCase := mock_usecases.NewMockTaskUseCase(controller)
	taskHandler := MakeTaskHandler(taskUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.TaskRoute+"/:id", authMiddleware.CheckAuth, taskHandler.GetSingleTask)
	}

	task := models.Task{Title: "title1", Description: "desc1"}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	taskUseCase.EXPECT().GetSingleTask(uint(11), uint(22)).Return(task, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.TaskRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newTask models.Task
	_ = json.Unmarshal(writer.Body.Bytes(), &newTask)
	assert.Equal(t, task, newTask)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.TaskRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCreateTask(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	taskUseCase := mock_usecases.NewMockTaskUseCase(controller)
	taskHandler := MakeTaskHandler(taskUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.BoardRoute+"/:id"+routes.ListRoute+"/:idL"+routes.TaskRoute, authMiddleware.CheckAuth, taskHandler.CreateTask)
	}
	task := models.Task{
		Title:       "title",
		Description: "desc",
		Position:    1,
		DateCreated: "22.02.02",
	}
	jsonNewTask, _ := json.Marshal(task)
	body := bytes.NewReader(jsonNewTask)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	taskUseCase.EXPECT().CreateTask(task, uint(11), uint(12), uint(22)).Return(&task, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/11"+routes.ListRoute+"/12"+routes.TaskRoute, body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newTask models.Task
	_ = json.Unmarshal(writer.Body.Bytes(), &newTask)
	assert.Equal(t, task, newTask)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/11"+routes.ListRoute+"/12"+routes.TaskRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRefactorTask(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	taskUseCase := mock_usecases.NewMockTaskUseCase(controller)
	taskHandler := MakeTaskHandler(taskUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.PUT(routes.TaskRoute+"/:id", authMiddleware.CheckAuth, taskHandler.RefactorTask)
	}
	task := models.Task{
		IdT:         11,
		Title:       "title",
		Description: "desc",
	}
	jsonNewTask, _ := json.Marshal(task)
	body := bytes.NewReader(jsonNewTask)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	taskUseCase.EXPECT().RefactorTask(task, uint(22)).Return(nil)
	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.TaskRoute+"/11", body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.TaskRoute+"/11", body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestDeleteTask(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	taskUseCase := mock_usecases.NewMockTaskUseCase(controller)
	taskHandler := MakeTaskHandler(taskUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.DELETE(routes.TaskRoute+"/:id", authMiddleware.CheckAuth, taskHandler.DeleteTask)
	}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	taskUseCase.EXPECT().DeleteTask(uint(11), uint(22)).Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.TaskRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.TaskRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
