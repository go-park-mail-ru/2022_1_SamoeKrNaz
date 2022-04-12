package handlers

import (
	mock_usecases "PLANEXA_backend/handlers/mocks"
	"PLANEXA_backend/middleware"
	mock_repositories "PLANEXA_backend/repositories/mocks"
	"PLANEXA_backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	boardUseCase := mock_usecases.NewMockBoardUseCase(controller)
	boardHandler := MakeBoardHandler(boardUseCase)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://planexa.netlify.app", "http://89.208.199.114:3000", "http://89.208.199.114:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	redis := mock_repositories.NewMockRedisRepository(controller)

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	boardUseCase.EXPECT().GetBoards(uint(22)).Return(nil, nil)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.GET("", middleware.CheckAuth, boardHandler.GetBoards)
	}

	request, _ := http.NewRequest("GET", routes.HomeRoute, nil)

	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

}
