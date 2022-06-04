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

func TestGetCheckLists(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListUseCase := mock_usecases.NewMockCheckListUseCase(controller)
	checkListHandler := MakeCheckListHandler(checkListUseCase)
	boardRepository, _, _ := CreateBoardMock()

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo, boardRepository)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.TaskRoute+"/:id", authMiddleware.CheckAuth, checkListHandler.GetCheckLists)
	}
	checkLists := []models.CheckList{{IdCl: 11, IdT: 11, Title: "title1", CheckListItems: nil}, {IdCl: 12, IdT: 12, Title: "title2", CheckListItems: nil}}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListUseCase.EXPECT().GetCheckLists(uint(22), uint(11)).Return(&checkLists, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.TaskRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newCheckLists []models.CheckList
	_ = json.Unmarshal(writer.Body.Bytes(), &newCheckLists)
	assert.Equal(t, checkLists, newCheckLists)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.TaskRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetCheckList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListUseCase := mock_usecases.NewMockCheckListUseCase(controller)
	checkListHandler := MakeCheckListHandler(checkListUseCase)
	boardRepository, _, _ := CreateBoardMock()

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo, boardRepository)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.CheckListRoute+"/:id", authMiddleware.CheckAuth, checkListHandler.GetSingleCheckList)
	}

	checkList := models.CheckList{IdCl: 11, IdT: 11, Title: "title1", CheckListItems: nil}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListUseCase.EXPECT().GetSingleCheckList(uint(11), uint(22)).Return(&checkList, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.CheckListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newCheckList models.CheckList
	_ = json.Unmarshal(writer.Body.Bytes(), &newCheckList)
	assert.Equal(t, checkList, newCheckList)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.CheckListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCreateCheckList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListUseCase := mock_usecases.NewMockCheckListUseCase(controller)
	checkListHandler := MakeCheckListHandler(checkListUseCase)
	boardRepository, _, _ := CreateBoardMock()

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo, boardRepository)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.TaskRoute+"/:id"+routes.CheckListRoute, authMiddleware.CheckAuth, checkListHandler.CreateCheckList)
	}
	checkList := models.CheckList{IdCl: 11, IdT: 11, Title: "title1", CheckListItems: nil}
	jsonNewCheckList, _ := json.Marshal(checkList)
	body := bytes.NewReader(jsonNewCheckList)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListUseCase.EXPECT().CreateCheckList(&checkList, uint(11), uint(22)).Return(&checkList, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.TaskRoute+"/11"+routes.CheckListRoute, body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newCheckList models.CheckList
	_ = json.Unmarshal(writer.Body.Bytes(), &newCheckList)
	assert.Equal(t, checkList, newCheckList)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.TaskRoute+"/11"+routes.CheckListRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRefactorCheckList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListUseCase := mock_usecases.NewMockCheckListUseCase(controller)
	checkListHandler := MakeCheckListHandler(checkListUseCase)
	boardRepository, _, _ := CreateBoardMock()

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo, boardRepository)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.PUT(routes.CheckListRoute+"/:id", authMiddleware.CheckAuth, checkListHandler.RefactorCheckList)
	}
	checkList := models.CheckList{IdCl: 11, IdT: 11, Title: "title1", CheckListItems: nil}
	jsonNewCheckList, _ := json.Marshal(checkList)
	body := bytes.NewReader(jsonNewCheckList)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListUseCase.EXPECT().RefactorCheckList(&checkList, uint(22)).Return(nil)
	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.CheckListRoute+"/11", body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.CheckListRoute+"/11", body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestDeleteCheckList(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	checkListUseCase := mock_usecases.NewMockCheckListUseCase(controller)
	checkListHandler := MakeCheckListHandler(checkListUseCase)
	boardRepository, _, _ := CreateBoardMock()

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo, boardRepository)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.DELETE(routes.CheckListRoute+"/:id", authMiddleware.CheckAuth, checkListHandler.DeleteCheckList)
	}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	checkListUseCase.EXPECT().DeleteCheckList(uint(11), uint(22)).Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.CheckListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.CheckListRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
