package handlers

import (
	"PLANEXA_backend/auth_microservice/server/handler"
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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLists(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	listUseCase := mock_usecases.NewMockListUseCase(controller)
	listHandler := MakeListHandler(listUseCase)

	router := gin.Default()

	redis := mock_repositories.NewMockRedisRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	grpcConn, _ := grpc.Dial(
		"session:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	sessService := handler.NewAuthCheckerClient(grpcConn)

	authMiddleware := middleware.CreateMiddleware(sessService)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.BoardRoute+"/:id", authMiddleware.CheckAuth, listHandler.GetLists)
	}
	tasks1 := []models.Task{{Title: "title1", Description: "desc1", DateCreated: "22.02.02"}, {Title: "title2", Description: "desc2", DateCreated: "23.02.02"}}
	tasks2 := []models.Task{{Title: "title3", Description: "desc3", DateCreated: "24.02.02"}, {Title: "title4", Description: "desc4", DateCreated: "25.02.02"}}
	lists := []models.List{{Title: "title1", Position: 1, Tasks: tasks1}, {Title: "title2", Position: 2, Tasks: tasks2}}

	//good
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	listUseCase.EXPECT().GetLists(uint(10), uint(22)).Return(lists, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.BoardRoute+"/10", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newLists []models.List
	_ = json.Unmarshal(writer.Body.Bytes(), &newLists)
	assert.Equal(t, lists, newLists)

	//bad
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.BoardRoute+"/10", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	listUseCase := mock_usecases.NewMockListUseCase(controller)
	listHandler := MakeListHandler(listUseCase)

	router := gin.Default()

	redis := mock_repositories.NewMockRedisRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	grpcConn, _ := grpc.Dial(
		"session:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	sessService := handler.NewAuthCheckerClient(grpcConn)

	authMiddleware := middleware.CreateMiddleware(sessService)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.ListRoute+"/:id", authMiddleware.CheckAuth, listHandler.GetSingleList)
	}

	tasks1 := []models.Task{{Title: "title1", Description: "desc1", DateCreated: "22.02.02"}, {Title: "title2", Description: "desc2", DateCreated: "23.02.02"}}
	list := models.List{Title: "title1", Position: 1, Tasks: tasks1}

	//good
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	listUseCase.EXPECT().GetSingleList(uint(11), uint(22)).Return(list, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.ListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newList models.List
	_ = json.Unmarshal(writer.Body.Bytes(), &newList)
	assert.Equal(t, list, newList)

	//bad
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.ListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCreateList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	listUseCase := mock_usecases.NewMockListUseCase(controller)
	listHandler := MakeListHandler(listUseCase)

	router := gin.Default()

	redis := mock_repositories.NewMockRedisRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	grpcConn, _ := grpc.Dial(
		"session:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	sessService := handler.NewAuthCheckerClient(grpcConn)

	authMiddleware := middleware.CreateMiddleware(sessService)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.BoardRoute+"/:id"+routes.ListRoute, authMiddleware.CheckAuth, listHandler.CreateList)
	}

	tasks1 := []models.Task{{Title: "title1", Description: "desc1", DateCreated: "22.02.02"}, {Title: "title2", Description: "desc2", DateCreated: "23.02.02"}}
	list := models.List{
		Title:    "title",
		Position: 1,
		Tasks:    tasks1,
	}
	jsonNewList, _ := json.Marshal(list)
	body := bytes.NewReader(jsonNewList)

	//good
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	listUseCase.EXPECT().CreateList(list, uint(11), uint(22)).Return(&list, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/11"+routes.ListRoute, body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newList models.List
	_ = json.Unmarshal(writer.Body.Bytes(), &newList)
	assert.Equal(t, list, newList)

	//bad
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/11"+routes.ListRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRefactorList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	listUseCase := mock_usecases.NewMockListUseCase(controller)
	listHandler := MakeListHandler(listUseCase)

	router := gin.Default()

	redis := mock_repositories.NewMockRedisRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	grpcConn, _ := grpc.Dial(
		"session:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	sessService := handler.NewAuthCheckerClient(grpcConn)

	authMiddleware := middleware.CreateMiddleware(sessService)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.PUT(routes.ListRoute+"/:id", authMiddleware.CheckAuth, listHandler.RefactorList)
	}
	tasks1 := []models.Task{{Title: "title1", Description: "desc1", DateCreated: "22.02.02"}, {Title: "title2", Description: "desc2", DateCreated: "23.02.02"}}
	list := models.List{
		IdL:      11,
		Title:    "title",
		Position: 1,
		Tasks:    tasks1,
	}
	jsonNewList, _ := json.Marshal(list)
	body := bytes.NewReader(jsonNewList)

	//good
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	listUseCase.EXPECT().RefactorList(list, uint(22), uint(11)).Return(nil)
	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.ListRoute+"/11", body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)

	//bad
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.ListRoute+"/11", body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestDeleteList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	listUseCase := mock_usecases.NewMockListUseCase(controller)
	listHandler := MakeListHandler(listUseCase)

	router := gin.Default()

	redis := mock_repositories.NewMockRedisRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	grpcConn, _ := grpc.Dial(
		"session:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	sessService := handler.NewAuthCheckerClient(grpcConn)

	authMiddleware := middleware.CreateMiddleware(sessService)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.DELETE(routes.ListRoute+"/:id", authMiddleware.CheckAuth, listHandler.DeleteList)
	}

	//good
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	listUseCase.EXPECT().DeleteList(uint(11), uint(22)).Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.ListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.ListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
