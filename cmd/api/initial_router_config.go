package main

import (
	handler_session "PLANEXA_backend/auth_microservice/server_session_ms/handler"
	customErrors "PLANEXA_backend/errors"
	handlers_impl "PLANEXA_backend/main_microservice/handlers"
	"PLANEXA_backend/main_microservice/middleware"
	repositories_impl "PLANEXA_backend/main_microservice/repositories/impl"
	usecases_impl "PLANEXA_backend/main_microservice/usecases/impl"
	"PLANEXA_backend/main_microservice/websocket/impl"
	"PLANEXA_backend/models"
	"PLANEXA_backend/routes"
	handler_user "PLANEXA_backend/user_microservice/server_user_ms/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	postgresHost     string
	postgresUser     string
	postgresPassword string
	postgresDbName   string
	postgresPort     string

	logFile string

	sessionContainer string
	userContainer    string

	metricsPath string
}

func ParseConfig() (conf Config) {
	viper.AddConfigPath("./cmd/api/")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	conf.postgresHost = viper.GetString("postgresHost")
	conf.postgresUser = viper.GetString("postgresUser")
	conf.postgresPassword = viper.GetString("postgresPassword")
	conf.postgresDbName = viper.GetString("postgresDbName")
	conf.postgresPort = viper.GetString("postgresPort")

	conf.logFile = viper.GetString("logFile")

	conf.sessionContainer = viper.GetString("sessionContainer")
	conf.userContainer = viper.GetString("userContainer")

	conf.metricsPath = viper.GetString("metricsPath")
	return
}

func initDB(conf Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(
		strings.Join([]string{"host=", conf.postgresHost, " user=", conf.postgresUser, " password=", conf.postgresPassword, " dbname=", conf.postgresDbName, " port=", conf.postgresPort}, "")))
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{}, &models.Board{}, &models.List{},
		&models.Task{}, &models.CheckList{}, &models.CheckListItem{}, &models.Comment{}, &models.Attachment{},
		&models.Notification{}, &models.ImportantTask{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initRouter() (*gin.Engine, error) {
	conf := ParseConfig()
	gin.DisableConsoleColor()
	f, err := os.Create(conf.logFile)
	if err != nil {
		return nil, err
	}

	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://planexa.ru", "http://planexa.netlify.app", "http://89.208.199.114:3000", "http://89.208.199.114:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true

	db, err := initDB(conf)
	if err != nil {
		return nil, err
	}
	grpcConn, err := grpc.Dial(
		conf.sessionContainer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, customErrors.ErrNoAccess
	}

	sessService := repositories_impl.CreateRepo(handler_session.NewAuthCheckerClient(grpcConn))

	if err != nil {
		return nil, err
	}
	grpcConnUser, err := grpc.Dial(
		conf.userContainer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, customErrors.ErrNoAccess
	}
	userService := repositories_impl.MakeUserRepository(db, handler_user.NewUserServiceClient(grpcConnUser))

	monitor := ginmetrics.GetMonitor()
	if err != nil {
		return nil, err
	}
	monitor.SetMetricPath(conf.metricsPath)
	monitor.SetSlowTime(10)
	monitor.Use(router)

	// ???????????????? ????????????????????????
	taskRepository := repositories_impl.MakeTaskRepository(db)
	listRepository := repositories_impl.MakeListRepository(db)
	boardRepository := repositories_impl.MakeBoardRepository(db)
	checkListRepository := repositories_impl.MakeCheckListRepository(db)
	checkListItemRepository := repositories_impl.MakeCheckListItemRepository(db)
	commentRepository := repositories_impl.MakeCommentRepository(db)
	attachmentRepository := repositories_impl.MakeAttachmentRepository(db)
	notificationRepository := repositories_impl.MakeNotificationRepository(db)

	webSocketPool := wsplanexa_impl.CreatePool()

	authMiddleware := middleware.CreateMiddleware(sessService, boardRepository, webSocketPool)

	router.Use(cors.New(config))

	router.GET(routes.HomeRoute+routes.WebSocketRoute, authMiddleware.CheckAuth, webSocketPool.Start)
	router.Use(middleware.CheckError())

	userHandler := handlers_impl.MakeUserHandler(usecases_impl.MakeUserUsecase(userService, sessService))
	taskHandler := handlers_impl.MakeTaskHandler(usecases_impl.MakeTaskUsecase(taskRepository, boardRepository, listRepository, userService, checkListRepository, commentRepository), usecases_impl.MakeNotificationUseCase(notificationRepository, boardRepository, taskRepository, userService))
	boardHandler := handlers_impl.MakeBoardHandler(usecases_impl.MakeBoardUsecase(boardRepository, listRepository, taskRepository, checkListRepository, userService, commentRepository), usecases_impl.MakeNotificationUseCase(notificationRepository, boardRepository, taskRepository, userService))
	listHandler := handlers_impl.MakeListHandler(usecases_impl.MakeListUsecase(listRepository, boardRepository))
	checkListHandler := handlers_impl.MakeCheckListHandler(usecases_impl.MakeCheckListUsecase(checkListRepository, taskRepository), usecases_impl.MakeTaskUsecase(taskRepository, boardRepository, listRepository, userService, checkListRepository, commentRepository))
	checkListItemHandler := handlers_impl.MakeCheckListItemHandler(usecases_impl.MakeCheckListItemUsecase(checkListItemRepository, checkListRepository, taskRepository), usecases_impl.MakeTaskUsecase(taskRepository, boardRepository, listRepository, userService, checkListRepository, commentRepository))
	commentHandler := handlers_impl.MakeCommentHandler(usecases_impl.MakeCommentUsecase(commentRepository, taskRepository, userService), usecases_impl.MakeTaskUsecase(taskRepository, boardRepository, listRepository, userService, checkListRepository, commentRepository))
	attachmentHandler := handlers_impl.MakeAttachmentHandler(usecases_impl.MakeAttachmentUseCase(attachmentRepository, taskRepository), usecases_impl.MakeTaskUsecase(taskRepository, boardRepository, listRepository, userService, checkListRepository, commentRepository))
	notificationHandler := handlers_impl.MakeNotificationHandler(usecases_impl.MakeNotificationUseCase(notificationRepository, boardRepository, taskRepository, userService))
	mainRoutes := router.Group(routes.HomeRoute)
	{
		boardRoutes := router.Group(routes.BoardRoute)
		{
			boardRoutes.POST("", authMiddleware.CheckAuth, boardHandler.CreateBoard)
			boardRoutes.PUT("/:id", authMiddleware.CheckAuth, boardHandler.RefactorBoard, authMiddleware.SendToWebSocket)
			boardRoutes.GET("/:id", authMiddleware.CheckAuth, boardHandler.GetSingleBoard)
			boardRoutes.DELETE("/:id", authMiddleware.CheckAuth, boardHandler.DeleteBoard, authMiddleware.SendToWebSocket)
			boardRoutes.POST("/:id/:idU", authMiddleware.CheckAuth, boardHandler.AppendUserToBoard, authMiddleware.SendToWebSocket, authMiddleware.SendNotification)
			boardRoutes.DELETE("/:id/:idU", authMiddleware.CheckAuth, boardHandler.DeleteUserToBoard, authMiddleware.SendToWebSocket, authMiddleware.SendNotification)
			boardRoutes.PUT("/:id/upload", authMiddleware.CheckAuth, boardHandler.SaveImage)
			boardRoutes.GET("/:id"+routes.ListRoute, authMiddleware.CheckAuth, listHandler.GetLists)
			boardRoutes.POST("/:id"+routes.ListRoute, authMiddleware.CheckAuth, listHandler.CreateList, authMiddleware.SendToWebSocket)
			boardRoutes.POST("/:id"+routes.ListRoute+"/:idL"+routes.TaskRoute, authMiddleware.CheckAuth, taskHandler.CreateTask, authMiddleware.SendToWebSocket)
			boardRoutes.POST("/append/:link", authMiddleware.CheckAuth, boardHandler.AppendUserToBoardByLink, authMiddleware.SendToWebSocket, authMiddleware.SendNotification)
		}
		listRoutes := router.Group(routes.ListRoute)
		{
			listRoutes.GET("/:id", authMiddleware.CheckAuth, listHandler.GetSingleList)
			listRoutes.PUT("/:id", authMiddleware.CheckAuth, listHandler.RefactorList, authMiddleware.SendToWebSocket)
			listRoutes.DELETE("/:id", authMiddleware.CheckAuth, listHandler.DeleteList, authMiddleware.SendToWebSocket)
			listRoutes.GET("/:id"+routes.TaskRoute, authMiddleware.CheckAuth, taskHandler.GetTasks)
		}
		taskRoutes := router.Group(routes.TaskRoute)
		{
			taskRoutes.GET("", authMiddleware.CheckAuth, taskHandler.GetImportantTasks)
			taskRoutes.GET("/:id", authMiddleware.CheckAuth, taskHandler.GetSingleTask)
			taskRoutes.PUT("/:id", authMiddleware.CheckAuth, taskHandler.RefactorTask, authMiddleware.SendToWebSocket)
			taskRoutes.DELETE("/:id", authMiddleware.CheckAuth, taskHandler.DeleteTask, authMiddleware.SendToWebSocket)
			taskRoutes.POST("/:id/:idU", authMiddleware.CheckAuth, taskHandler.AppendUserToTask, authMiddleware.SendToWebSocket, authMiddleware.SendNotification)
			taskRoutes.DELETE("/:id/:idU", authMiddleware.CheckAuth, taskHandler.DeleteUserFromTask, authMiddleware.SendToWebSocket, authMiddleware.SendNotification)
			taskRoutes.GET("/:id"+routes.CheckListRoute, authMiddleware.CheckAuth, checkListHandler.GetCheckLists)
			taskRoutes.POST("/:id"+routes.CheckListRoute, authMiddleware.CheckAuth, checkListHandler.CreateCheckList, authMiddleware.SendToWebSocket)
			taskRoutes.GET("/:id"+routes.CommentRoute, authMiddleware.CheckAuth, commentHandler.GetComments)
			taskRoutes.POST("/:id"+routes.CommentRoute, authMiddleware.CheckAuth, commentHandler.CreateComment, authMiddleware.SendToWebSocket)
			taskRoutes.PUT("/:id"+routes.AttachmentRoute, authMiddleware.CheckAuth, attachmentHandler.CreateAttachment, authMiddleware.SendToWebSocket)
			taskRoutes.POST("/append/:link", authMiddleware.CheckAuth, taskHandler.AppendUserToTaskByLink, authMiddleware.SendToWebSocket)
		}
		checkListRoutes := router.Group(routes.CheckListRoute)
		{
			checkListRoutes.GET("/:id", authMiddleware.CheckAuth, checkListHandler.GetSingleCheckList)
			checkListRoutes.PUT("/:id", authMiddleware.CheckAuth, checkListHandler.RefactorCheckList, authMiddleware.SendToWebSocket)
			checkListRoutes.DELETE("/:id", authMiddleware.CheckAuth, checkListHandler.DeleteCheckList, authMiddleware.SendToWebSocket)
			checkListRoutes.GET("/:id"+routes.CheckListItemRoute, authMiddleware.CheckAuth, checkListItemHandler.GetCheckListItems)
			checkListRoutes.POST("/:id"+routes.CheckListItemRoute, authMiddleware.CheckAuth, checkListItemHandler.CreateCheckListItem, authMiddleware.SendToWebSocket)
		}
		checkListItemRoutes := router.Group(routes.CheckListItemRoute)
		{
			checkListItemRoutes.GET("/:id", authMiddleware.CheckAuth, checkListItemHandler.GetSingleCheckListItem)
			checkListItemRoutes.PUT("/:id", authMiddleware.CheckAuth, checkListItemHandler.RefactorCheckListItem, authMiddleware.SendToWebSocket)
			checkListItemRoutes.DELETE("/:id", authMiddleware.CheckAuth, checkListItemHandler.DeleteCheckListItem, authMiddleware.SendToWebSocket)
		}
		commentRoutes := router.Group(routes.CommentRoute)
		{
			commentRoutes.GET("/:id", authMiddleware.CheckAuth, commentHandler.GetSingleComment)
			commentRoutes.PUT("/:id", authMiddleware.CheckAuth, commentHandler.RefactorComment, authMiddleware.SendToWebSocket)
			commentRoutes.DELETE("/:id", authMiddleware.CheckAuth, commentHandler.DeleteComment, authMiddleware.SendToWebSocket)
		}
		attachmentRoutes := router.Group(routes.AttachmentRoute)
		{
			attachmentRoutes.GET("/:id", authMiddleware.CheckAuth, attachmentHandler.GetSingleAttachment)
			attachmentRoutes.DELETE("/:id", authMiddleware.CheckAuth, attachmentHandler.DeleteAttachment, authMiddleware.SendToWebSocket)
		}
		notificationRoutes := router.Group(routes.NotificationRoute)
		{
			notificationRoutes.GET("", authMiddleware.CheckAuth, notificationHandler.GetNotifications)
			notificationRoutes.POST("", authMiddleware.CheckAuth, notificationHandler.ReadNotifications)
			notificationRoutes.DELETE("", authMiddleware.CheckAuth, notificationHandler.DeleteNotifications)
		}
		mainRoutes.POST(routes.LoginRoute, userHandler.Login)
		mainRoutes.GET("/get"+routes.BoardRoute+"s", authMiddleware.CheckAuth, boardHandler.GetBoards)
		mainRoutes.POST(routes.RegisterRoute, userHandler.Register)
		mainRoutes.DELETE(routes.LogoutRoute, userHandler.Logout)
		mainRoutes.GET(routes.ProfileRoute+"/:id", authMiddleware.CheckAuth, userHandler.GetInfoById)
		mainRoutes.GET(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.GetInfoByCookie)
		mainRoutes.PUT(routes.ProfileRoute+"/upload", authMiddleware.CheckAuth, userHandler.SaveAvatar)
		mainRoutes.PUT(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.RefactorProfile)
		mainRoutes.POST(routes.ProfileRoute+"/like", authMiddleware.CheckAuth, userHandler.GetUsersLike)
	}
	return router, nil
}
