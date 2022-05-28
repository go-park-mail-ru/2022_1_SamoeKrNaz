package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/middleware"
	mock_repositories "PLANEXA_backend/main_microservice/repositories/mocks"
	mock_usecases "PLANEXA_backend/main_microservice/usecases/mocks"
	"PLANEXA_backend/models"
	"PLANEXA_backend/routes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAttachment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	attachmentUseCase := mock_usecases.NewMockAttachmentUseCase(controller)
	attachmentHandler := MakeAttachmentHandler(attachmentUseCase)
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
		mainRoutes.GET(routes.AttachmentRoute+"/:id", authMiddleware.CheckAuth, attachmentHandler.GetSingleAttachment)
	}

	attachment := models.Attachment{IdA: 1}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	attachmentUseCase.EXPECT().GetById(uint(11), uint(22)).Return(&attachment, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.AttachmentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newAttachment models.Attachment
	_ = json.Unmarshal(writer.Body.Bytes(), &newAttachment)
	assert.Equal(t, attachment, newAttachment)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.AttachmentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestDeleteAttachment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	attachmentUseCase := mock_usecases.NewMockAttachmentUseCase(controller)
	attachmentHandler := MakeAttachmentHandler(attachmentUseCase)
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
		mainRoutes.DELETE(routes.AttachmentRoute+"/:id", authMiddleware.CheckAuth, attachmentHandler.DeleteAttachment)
	}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	attachmentUseCase.EXPECT().DeleteAttachment(uint(11), uint(22)).Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.AttachmentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.AttachmentRoute+"/11", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
