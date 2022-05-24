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

func TestGetBoards(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET("", authMiddleware.CheckAuth, boardHandler.GetBoards)
	}
	boards := []models.Board{{Title: "title1", Description: "desc1"}, {Title: "title2", Description: "desc2"}}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().GetBoards(uint(22)).Return(boards, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute, nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newBoards []models.Board
	_ = json.Unmarshal(writer.Body.Bytes(), &newBoards)
	assert.Equal(t, boards, newBoards)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute, nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.BoardRoute+"/:id", authMiddleware.CheckAuth, boardHandler.GetSingleBoard)
	}

	board := models.Board{Title: "title1", Description: "desc1"}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().GetBoard(uint(11), uint(22)).Return(board, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.BoardRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newBoard models.Board
	_ = json.Unmarshal(writer.Body.Bytes(), &newBoard)
	assert.Equal(t, board, newBoard)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.BoardRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCreateBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.BoardRoute, authMiddleware.CheckAuth, boardHandler.CreateBoard)
	}
	board := models.Board{
		Title:       "title",
		Description: "desc",
	}
	jsonNewBoard, _ := json.Marshal(board)
	body := bytes.NewReader(jsonNewBoard)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().CreateBoard(uint(22), board).Return(&models.Board{Title: "title", Description: "desc"}, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute, body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)
	var newBoard models.Board
	_ = json.Unmarshal(writer.Body.Bytes(), &newBoard)
	assert.Equal(t, board, newBoard)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRefactorBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.PUT(routes.BoardRoute+"/:id", authMiddleware.CheckAuth, boardHandler.RefactorBoard)
	}
	board := models.Board{
		IdB:         11,
		Title:       "title",
		Description: "desc",
	}
	jsonNewBoard, _ := json.Marshal(board)
	body := bytes.NewReader(jsonNewBoard)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().RefactorBoard(uint(22), board).Return(nil)
	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.BoardRoute+"/11", body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.BoardRoute+"/11", body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestDeleteBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.DELETE(routes.BoardRoute+"/:id", authMiddleware.CheckAuth, boardHandler.DeleteBoard)
	}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().DeleteBoard(uint(11), uint(22)).Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.BoardRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.BoardRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestAppendUserToBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.BoardRoute+"/:id/:idU", authMiddleware.CheckAuth, boardHandler.AppendUserToBoard)
	}

	var user models.User

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().AppendUserToBoard(uint(22), uint(15), uint(11)).Return(user, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/11/15", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/11/15", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestAppendUserToBoardLinkBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.BoardRoute+"/append/:link", authMiddleware.CheckAuth, boardHandler.AppendUserToBoardByLink)
	}

	var board models.Board

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().AppendUserByLink(uint(22), "link").Return(&board, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/append/link", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.BoardRoute+"/append/link", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
