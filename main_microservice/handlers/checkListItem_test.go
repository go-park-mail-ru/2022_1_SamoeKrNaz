package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/middleware"
	"PLANEXA_backend/main_microservice/repositories/mocks"
	"PLANEXA_backend/main_microservice/usecases/mocks"
	"PLANEXA_backend/models"
	"PLANEXA_backend/routes"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCheckListItems(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListItemUseCase := mock_usecases.NewMockCheckListItemUseCase(controller)
	checkListItemHandler := MakeCheckListItemHandler(checkListItemUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.CheckListRoute+"/:id"+routes.CheckListItemRoute, authMiddleware.CheckAuth, checkListItemHandler.GetCheckListItems)
	}
	checkListItems := []models.CheckListItem{{IdT: 11, IdCl: 11, IdClIt: 11, IsReady: true, Description: "desc1"},
		{IdT: 11, IdCl: 11, IdClIt: 11, IsReady: true, Description: "desc2"}}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListItemUseCase.EXPECT().GetCheckListItems(uint(22), uint(11)).Return(&checkListItems, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.CheckListRoute+"/11"+routes.CheckListItemRoute, nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newCheckListItems []models.CheckListItem
	_ = json.Unmarshal(writer.Body.Bytes(), &newCheckListItems)
	assert.Equal(t, checkListItems, newCheckListItems)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.CheckListRoute+"/11"+routes.CheckListItemRoute, nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetCheckListItem(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListItemUseCase := mock_usecases.NewMockCheckListItemUseCase(controller)
	checkListItemHandler := MakeCheckListItemHandler(checkListItemUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.CheckListItemRoute+"/:id", authMiddleware.CheckAuth, checkListItemHandler.GetSingleCheckListItem)
	}

	checkListItem := models.CheckListItem{IdT: 11, IdCl: 11, IdClIt: 11, IsReady: true, Description: "desc1"}
	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListItemUseCase.EXPECT().GetSingleCheckListItem(uint(11), uint(22)).Return(&checkListItem, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.CheckListItemRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newCheckListItem models.CheckListItem
	_ = json.Unmarshal(writer.Body.Bytes(), &newCheckListItem)
	assert.Equal(t, checkListItem, newCheckListItem)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.CheckListItemRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCreateCheckListItem(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListItemUseCase := mock_usecases.NewMockCheckListItemUseCase(controller)
	checkListItemHandler := MakeCheckListItemHandler(checkListItemUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.CheckListRoute+"/:id"+routes.CheckListItemRoute, authMiddleware.CheckAuth, checkListItemHandler.CreateCheckListItem)
	}
	checkListItem := models.CheckListItem{IdT: 11, IdCl: 11, IdClIt: 11, IsReady: true, Description: "desc1"}
	jsonNewCheckListItem, _ := json.Marshal(checkListItem)
	body := bytes.NewReader(jsonNewCheckListItem)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListItemUseCase.EXPECT().CreateCheckListItem(&checkListItem, uint(11), uint(22)).Return(&checkListItem, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.CheckListRoute+"/11"+routes.CheckListItemRoute, body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newCheckListItem models.CheckListItem
	_ = json.Unmarshal(writer.Body.Bytes(), &newCheckListItem)
	assert.Equal(t, checkListItem, newCheckListItem)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.CheckListRoute+"/11"+routes.CheckListItemRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRefactorCheckListItem(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListItemUseCase := mock_usecases.NewMockCheckListItemUseCase(controller)
	checkListItemHandler := MakeCheckListItemHandler(checkListItemUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.PUT(routes.CheckListItemRoute+"/:id", authMiddleware.CheckAuth, checkListItemHandler.RefactorCheckListItem)
	}
	checkListItem := models.CheckListItem{IdT: 11, IdCl: 11, IdClIt: 11, IsReady: true, Description: "desc1"}
	jsonNewCheckListItem, _ := json.Marshal(checkListItem)
	body := bytes.NewReader(jsonNewCheckListItem)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListItemUseCase.EXPECT().RefactorCheckListItem(&checkListItem, uint(22)).Return(nil)
	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.CheckListItemRoute+"/11", body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.CheckListItemRoute+"/11", body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestDeleteCheckListItem(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListItemUseCase := mock_usecases.NewMockCheckListItemUseCase(controller)
	checkListItemHandler := MakeCheckListItemHandler(checkListItemUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.DELETE(routes.CheckListItemRoute+"/:id", authMiddleware.CheckAuth, checkListItemHandler.DeleteCheckListItem)
	}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListItemUseCase.EXPECT().DeleteCheckListItem(uint(11), uint(22)).Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.CheckListItemRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.CheckListItemRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
