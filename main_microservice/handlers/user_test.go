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

func TestLogin(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	userUseCase := mock_usecases.NewMockUserUseCase(controller)
	userHandler := MakeUserHandler(userUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.LoginRoute, userHandler.Login)
	}
	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
	jsonUser, _ := json.Marshal(user)
	body := bytes.NewReader(jsonUser)

	//good
	userUseCase.EXPECT().Login(user).Return(uint(22), "token", nil)
	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.LoginRoute, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newUser models.User
	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
	assert.Equal(t, user, newUser)

	//bad
	body = bytes.NewReader(jsonUser)
	userUseCase.EXPECT().Login(user).Return(uint(0), "", customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.LoginRoute, body)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRegister(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	userUseCase := mock_usecases.NewMockUserUseCase(controller)
	userHandler := MakeUserHandler(userUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.RegisterRoute, userHandler.Register)
	}

	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
	jsonUser, _ := json.Marshal(user)
	body := bytes.NewReader(jsonUser)

	//good
	userUseCase.EXPECT().Register(user).Return(uint(22), "token", nil)
	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.RegisterRoute, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)
	var newUser models.User
	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
	assert.Equal(t, user, newUser)

	//bad
	body = bytes.NewReader(jsonUser)
	userUseCase.EXPECT().Register(user).Return(uint(0), "", customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.RegisterRoute, body)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestLogout(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	userUseCase := mock_usecases.NewMockUserUseCase(controller)
	userHandler := MakeUserHandler(userUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.DELETE(routes.LogoutRoute, userHandler.Logout)
	}

	//good
	userUseCase.EXPECT().Logout("sess1").Return(nil)
	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.LogoutRoute, nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	userUseCase.EXPECT().Logout("sess1").Return(customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.LogoutRoute, nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetInfoById(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	userUseCase := mock_usecases.NewMockUserUseCase(controller)
	userHandler := MakeUserHandler(userUseCase)
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
		mainRoutes.GET(routes.ProfileRoute+"/:id", authMiddleware.CheckAuth, userHandler.GetInfoById)
	}

	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute+"/22", nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newUser models.User
	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
	assert.Equal(t, user, newUser)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute+"/22", nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetInfoByCookie(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	userUseCase := mock_usecases.NewMockUserUseCase(controller)
	userHandler := MakeUserHandler(userUseCase)
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
		mainRoutes.GET(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.GetInfoByCookie)
	}

	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute, nil)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var newUser models.User
	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
	assert.Equal(t, user, newUser)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute, nil)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestRefactorProfile(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	userUseCase := mock_usecases.NewMockUserUseCase(controller)
	userHandler := MakeUserHandler(userUseCase)
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
		mainRoutes.PUT(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.RefactorProfile)
	}

	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
	jsonNewUser, _ := json.Marshal(user)
	body := bytes.NewReader(jsonNewUser)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	userUseCase.EXPECT().RefactorProfile(user).Return(nil)
	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.ProfileRoute, body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	body = bytes.NewReader(jsonNewUser)
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.ProfileRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)

	body = bytes.NewReader(jsonNewUser)
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	userUseCase.EXPECT().RefactorProfile(user).Return(customErrors.ErrUsernameExist)
	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.ProfileRoute, body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusConflict, writer.Code)
}

func TestGetUsersLike(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	userUseCase := mock_usecases.NewMockUserUseCase(controller)
	userHandler := MakeUserHandler(userUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())
	sessionRepo := mock_repositories.NewMockSessionRepository(controller)
	boardRepository, _, _ := CreateBoardMock()

	cookie := &http.Cookie{
		Name:  "token",
		Value: "sess1",
	}

	authMiddleware := middleware.CreateMiddleware(sessionRepo, boardRepository)

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.ProfileRoute+"/like", authMiddleware.CheckAuth, userHandler.GetUsersLike)
	}

	var users []models.User
	user := models.User{Username: "user1"}
	jsonUser, _ := json.Marshal(user)
	body := bytes.NewReader(jsonUser)

	//good
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
	userUseCase.EXPECT().GetUsersLike(user.Username).Return(&users, nil)
	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.ProfileRoute+"/like", body)
	request.AddCookie(cookie)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	//bad
	sessionRepo.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.ProfileRoute+"/like", body)
	request.AddCookie(cookie)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
