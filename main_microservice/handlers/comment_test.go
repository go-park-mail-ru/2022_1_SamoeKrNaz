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

func TestGetComments(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	commentUseCase := mock_usecases.NewMockCommentUseCase(controller)
	commentHandler := MakeCommentHandler(commentUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.TaskRoute+"/:id"+routes.CommentRoute, authMiddleware.CheckAuth, commentHandler.GetComments)
	}
	comments := []models.Comment{{IdCm: 11, Text: "text1", DateCreated: "01.01.01", IdT: 11, IdU: 22, User: models.User{}},
		{IdCm: 11, Text: "text2", DateCreated: "01.01.02", IdT: 11, IdU: 22, User: models.User{}}}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	commentUseCase.EXPECT().GetComments(uint(22), uint(11)).Return(&comments, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.TaskRoute+"/11"+routes.CommentRoute, nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newComments []models.Comment
	_ = json.Unmarshal(writer.Body.Bytes(), &newComments)
	assert.Equal(t, comments, newComments)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.TaskRoute+"/11"+routes.CommentRoute, nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetComment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	commentUseCase := mock_usecases.NewMockCommentUseCase(controller)
	commentHandler := MakeCommentHandler(commentUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET(routes.CommentRoute+"/:id", authMiddleware.CheckAuth, commentHandler.GetSingleComment)
	}

	comment := models.Comment{IdCm: 11, Text: "text1", DateCreated: "01.01.01", IdT: 11, IdU: 22, User: models.User{}}
	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	commentUseCase.EXPECT().GetSingleComment(uint(11), uint(22)).Return(&comment, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.CommentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newComment models.Comment
	_ = json.Unmarshal(writer.Body.Bytes(), &newComment)
	assert.Equal(t, comment, newComment)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.CommentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCreateComment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	commentUseCase := mock_usecases.NewMockCommentUseCase(controller)
	commentHandler := MakeCommentHandler(commentUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.TaskRoute+"/:id"+routes.CommentRoute, authMiddleware.CheckAuth, commentHandler.CreateComment)
	}
	comment := models.Comment{IdCm: 11, Text: "text1", DateCreated: "01.01.01", IdT: 11, IdU: 22, User: models.User{}}
	jsonNewComment, _ := json.Marshal(comment)
	body := bytes.NewReader(jsonNewComment)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	commentUseCase.EXPECT().CreateComment(&comment, uint(11), uint(22)).Return(&comment, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.TaskRoute+"/11"+routes.CommentRoute, body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newComment models.Comment
	_ = json.Unmarshal(writer.Body.Bytes(), &newComment)
	assert.Equal(t, comment, newComment)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.TaskRoute+"/11"+routes.CommentRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRefactorComment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	commentUseCase := mock_usecases.NewMockCommentUseCase(controller)
	commentHandler := MakeCommentHandler(commentUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.PUT(routes.CommentRoute+"/:id", authMiddleware.CheckAuth, commentHandler.RefactorComment)
	}
	comment := models.Comment{IdCm: 11, Text: "text1", DateCreated: "01.01.01", IdT: 11, IdU: 22, User: models.User{}}
	jsonNewComment, _ := json.Marshal(comment)
	body := bytes.NewReader(jsonNewComment)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	commentUseCase.EXPECT().RefactorComment(&comment, uint(22)).Return(nil)
	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.CommentRoute+"/11", body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.CommentRoute+"/11", body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestDeleteComment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	commentUseCase := mock_usecases.NewMockCommentUseCase(controller)
	commentHandler := MakeCommentHandler(commentUseCase)

	router := gin.Default()

	sessionRepo := mock_repositories.NewMockSessionRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.DELETE(routes.CommentRoute+"/:id", authMiddleware.CheckAuth, commentHandler.DeleteComment)
	}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	commentUseCase.EXPECT().DeleteComment(uint(11), uint(22)).Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.CommentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.CommentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
