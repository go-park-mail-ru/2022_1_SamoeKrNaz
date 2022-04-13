package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/middleware"
	mock_repositories "PLANEXA_backend/repositories/mocks"
	"PLANEXA_backend/routes"
	"PLANEXA_backend/usecases/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()

	redis := mock_repositories.NewMockRedisRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(redis)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET("", authMiddleware.CheckAuth, boardHandler.GetBoards)
	}

	//good
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().GetBoards(uint(22)).Return(nil, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute, nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute, nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
