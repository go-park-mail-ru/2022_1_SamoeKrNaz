package handlers

//
//import (
//	"PLANEXA_backend/auth_microservice/server/handler"
//	customErrors "PLANEXA_backend/errors"
//	"PLANEXA_backend/middleware"
//	"PLANEXA_backend/models"
//	mock_repositories "PLANEXA_backend/repositories/mocks"
//	"PLANEXA_backend/routes"
//	"PLANEXA_backend/usecases/mocks"
//	"bytes"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestLogin(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	userUseCase := mock_usecases.NewMockUserUseCase(controller)
//	userHandler := MakeUserHandler(userUseCase)
//
//	router := gin.Default()
//
//	mainRoutes := router.Group(routes.HomeRoute)
//	{
//		mainRoutes.POST(routes.LoginRoute, userHandler.Login)
//	}
//	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
//	jsonUser, _ := json.Marshal(user)
//	body := bytes.NewReader(jsonUser)
//
//	//good
//	userUseCase.EXPECT().Login(user).Return(uint(22), "token", nil)
//	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
//	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.LoginRoute, body)
//	writer := httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusOK, writer.Code)
//	var newUser models.User
//	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
//	assert.Equal(t, user, newUser)
//
//	//bad
//	body = bytes.NewReader(jsonUser)
//	userUseCase.EXPECT().Login(user).Return(uint(0), "", customErrors.ErrUnauthorized)
//	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.LoginRoute, body)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusUnauthorized, writer.Code)
//}
//
//func TestRegister(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	userUseCase := mock_usecases.NewMockUserUseCase(controller)
//	userHandler := MakeUserHandler(userUseCase)
//
//	router := gin.Default()
//
//	mainRoutes := router.Group(routes.HomeRoute)
//	{
//		mainRoutes.POST(routes.RegisterRoute, userHandler.Register)
//	}
//
//	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
//	jsonUser, _ := json.Marshal(user)
//	body := bytes.NewReader(jsonUser)
//
//	//good
//	userUseCase.EXPECT().Register(user).Return(uint(22), "token", nil)
//	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
//	request, _ := http.NewRequest("POST", routes.HomeRoute+routes.RegisterRoute, body)
//	writer := httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusCreated, writer.Code)
//	var newUser models.User
//	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
//	assert.Equal(t, user, newUser)
//
//	//bad
//	body = bytes.NewReader(jsonUser)
//	userUseCase.EXPECT().Register(user).Return(uint(0), "", customErrors.ErrUnauthorized)
//	request, _ = http.NewRequest("POST", routes.HomeRoute+routes.RegisterRoute, body)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusUnauthorized, writer.Code)
//}
//
//func TestLogout(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	userUseCase := mock_usecases.NewMockUserUseCase(controller)
//	userHandler := MakeUserHandler(userUseCase)
//
//	router := gin.Default()
//
//	cookie := &http.Cookie{
//		Name:  "token",
//		Value: "sess1",
//	}
//
//	mainRoutes := router.Group(routes.HomeRoute)
//	{
//		mainRoutes.DELETE(routes.LogoutRoute, userHandler.Logout)
//	}
//
//	//good
//	userUseCase.EXPECT().Logout("sess1").Return(nil)
//	request, _ := http.NewRequest("DELETE", routes.HomeRoute+routes.LogoutRoute, nil)
//	request.AddCookie(cookie)
//	writer := httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusOK, writer.Code)
//
//	//bad
//	userUseCase.EXPECT().Logout("sess1").Return(customErrors.ErrUnauthorized)
//	request, _ = http.NewRequest("DELETE", routes.HomeRoute+routes.LogoutRoute, nil)
//	request.AddCookie(cookie)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusUnauthorized, writer.Code)
//}
//
//func TestGetInfoById(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	userUseCase := mock_usecases.NewMockUserUseCase(controller)
//	userHandler := MakeUserHandler(userUseCase)
//
//	router := gin.Default()
//	redis := mock_repositories.NewMockRedisRepository(controller)
//
//	cookie := &http.Cookie{
//		Name:  "token",
//		Value: "sess1",
//	}
//	grpcConn, _ := grpc.Dial(
//		"session:8081",
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//
//	sessService := handler.NewAuthCheckerClient(grpcConn)
//
//	authMiddleware := middleware.CreateMiddleware(sessService)
//
//	mainRoutes := router.Group(routes.HomeRoute)
//	{
//		mainRoutes.GET(routes.ProfileRoute+"/:id", authMiddleware.CheckAuth, userHandler.GetInfoById)
//	}
//
//	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
//
//	//good
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
//	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
//	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute+"/22", nil)
//	request.AddCookie(cookie)
//	writer := httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusOK, writer.Code)
//	var newUser models.User
//	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
//	assert.Equal(t, user, newUser)
//
//	//bad
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
//	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute+"/22", nil)
//	request.AddCookie(cookie)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusUnauthorized, writer.Code)
//}
//
//func TestGetInfoByCookie(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	userUseCase := mock_usecases.NewMockUserUseCase(controller)
//	userHandler := MakeUserHandler(userUseCase)
//
//	router := gin.Default()
//	redis := mock_repositories.NewMockRedisRepository(controller)
//
//	cookie := &http.Cookie{
//		Name:  "token",
//		Value: "sess1",
//	}
//	grpcConn, _ := grpc.Dial(
//		"session:8081",
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//
//	sessService := handler.NewAuthCheckerClient(grpcConn)
//
//	authMiddleware := middleware.CreateMiddleware(sessService)
//
//	mainRoutes := router.Group(routes.HomeRoute)
//	{
//		mainRoutes.GET(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.GetInfoByCookie)
//	}
//
//	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
//
//	//good
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
//	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
//	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute, nil)
//	request.AddCookie(cookie)
//	writer := httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusOK, writer.Code)
//	var newUser models.User
//	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
//	assert.Equal(t, user, newUser)
//
//	//bad
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
//	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute, nil)
//	request.AddCookie(cookie)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusUnauthorized, writer.Code)
//}
//
///*
//func TestSaveAvatar(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	userUseCase := mock_usecases.NewMockUserUseCase(controller)
//	userHandler := MakeUserHandler(userUseCase)
//
//	router := gin.Default()
//	redis := mock_repositories.NewMockRedisRepository(controller)
//
//	cookie := &http.Cookie{
//		Name:  "token",
//		Value: "sess1",
//	}
//	authMiddleware := middleware.CreateMiddleware(redis)
//
//	mainRoutes := router.Group(routes.HomeRoute)
//	{
//		mainRoutes.GET(routes.ProfileRoute+"/upload", authMiddleware.CheckAuth, userHandler.SaveAvatar)
//	}
//
//	//good
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
//	userUseCase.EXPECT().GetInfoById(uint(22)).Return(user, nil)
//	request, _ := http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute+"/upload", nil)
//	request.AddCookie(cookie)
//	writer := httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusOK, writer.Code)
//	var newUser models.User
//	_ = json.Unmarshal(writer.Body.Bytes(), &newUser)
//	assert.Equal(t, user, newUser)
//
//	//bad
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
//	request, _ = http.NewRequest("GET", routes.HomeRoute+routes.ProfileRoute+"/upload", nil)
//	request.AddCookie(cookie)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusUnauthorized, writer.Code)
//}*/
//
//func TestRefactorProfile(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//	userUseCase := mock_usecases.NewMockUserUseCase(controller)
//	userHandler := MakeUserHandler(userUseCase)
//
//	router := gin.Default()
//	redis := mock_repositories.NewMockRedisRepository(controller)
//
//	cookie := &http.Cookie{
//		Name:  "token",
//		Value: "sess1",
//	}
//	grpcConn, _ := grpc.Dial(
//		"session:8081",
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//
//	sessService := handler.NewAuthCheckerClient(grpcConn)
//
//	authMiddleware := middleware.CreateMiddleware(sessService)
//
//	mainRoutes := router.Group(routes.HomeRoute)
//	{
//		mainRoutes.PUT(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.RefactorProfile)
//	}
//
//	user := models.User{IdU: 22, Username: "user1", Password: "pass1"}
//	jsonNewUser, _ := json.Marshal(user)
//	body := bytes.NewReader(jsonNewUser)
//
//	//good
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
//	userUseCase.EXPECT().RefactorProfile(user).Return(nil)
//	request, _ := http.NewRequest("PUT", routes.HomeRoute+routes.ProfileRoute, body)
//	request.AddCookie(cookie)
//	writer := httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusOK, writer.Code)
//
//	//bad
//	body = bytes.NewReader(jsonNewUser)
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(0), customErrors.ErrUnauthorized)
//	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.ProfileRoute, body)
//	request.AddCookie(cookie)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusUnauthorized, writer.Code)
//
//	body = bytes.NewReader(jsonNewUser)
//	redis.EXPECT().GetSession(cookie.Value).Return(uint64(22), nil)
//	userUseCase.EXPECT().RefactorProfile(user).Return(customErrors.ErrUsernameExist)
//	request, _ = http.NewRequest("PUT", routes.HomeRoute+routes.ProfileRoute, body)
//	request.AddCookie(cookie)
//	writer = httptest.NewRecorder()
//	router.ServeHTTP(writer, request)
//	assert.Equal(t, http.StatusBadRequest, writer.Code)
//}
