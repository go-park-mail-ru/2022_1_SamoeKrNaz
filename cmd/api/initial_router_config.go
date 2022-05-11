package main

import (
	handler_session "PLANEXA_backend/auth_microservice/server_session_ms/handler"
	customErrors "PLANEXA_backend/errors"
	handlers_impl "PLANEXA_backend/main_microservice/handlers"
	"PLANEXA_backend/main_microservice/middleware"
	repositories_impl "PLANEXA_backend/main_microservice/repositories/impl"
	usecases_impl "PLANEXA_backend/main_microservice/usecases/impl"
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
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	conf.postgresHost = viper.GetString("postgresHost")
	conf.postgresUser = viper.GetString("postgresUser")
	conf.postgresPassword = viper.GetString("postgresPassword")
	conf.postgresDbName = viper.GetString("postgresDbName")
	conf.postgresPort = viper.GetString("postgresPort")

	conf.logFile = viper.GetString("gin.log")

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
		&models.Task{}, &models.CheckList{}, &models.CheckListItem{}, &models.Comment{})
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
	config.AllowOrigins = []string{"http://localhost:3000", "http://planexa.netlify.app", "http://89.208.199.114:3000", "http://89.208.199.114:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

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

	// создание репозиториев
	taskRepository := repositories_impl.MakeTaskRepository(db)
	listRepository := repositories_impl.MakeListRepository(db)
	boardRepository := repositories_impl.MakeBoardRepository(db)
	checkListRepository := repositories_impl.MakeCheckListRepository(db)
	checkListItemRepository := repositories_impl.MakeCheckListItemRepository(db)
	commentRepository := repositories_impl.MakeCommentRepository(db)
	attachmentRepository := repositories_impl.MakeAttachmentRepository(db)

	authMiddleware := middleware.CreateMiddleware(sessService)

	userHandler := handlers_impl.MakeUserHandler(usecases_impl.MakeUserUsecase(userService, sessService))
	taskHandler := handlers_impl.MakeTaskHandler(usecases_impl.MakeTaskUsecase(taskRepository, boardRepository, listRepository, userService, checkListRepository, commentRepository))
	boardHandler := handlers_impl.MakeBoardHandler(usecases_impl.MakeBoardUsecase(boardRepository, listRepository, taskRepository, checkListRepository, userService, commentRepository))
	listHandler := handlers_impl.MakeListHandler(usecases_impl.MakeListUsecase(listRepository, boardRepository))
	checkListHandler := handlers_impl.MakeCheckListHandler(usecases_impl.MakeCheckListUsecase(checkListRepository, taskRepository))
	checkListItemHandler := handlers_impl.MakeCheckListItemHandler(usecases_impl.MakeCheckListItemUsecase(checkListItemRepository, checkListRepository, taskRepository))
	commentHandler := handlers_impl.MakeCommentHandler(usecases_impl.MakeCommentUsecase(commentRepository, taskRepository, userService))
	attachmentHandler := handlers_impl.MakeAttachmentHandler(usecases_impl.MakeAttachmentUseCase(attachmentRepository, taskRepository))
	mainRoutes := router.Group(routes.HomeRoute)
	{
		boardRoutes := router.Group(routes.BoardRoute)
		{
			boardRoutes.POST("", authMiddleware.CheckAuth, boardHandler.CreateBoard)
			boardRoutes.PUT("/:id", authMiddleware.CheckAuth, boardHandler.RefactorBoard)
			boardRoutes.GET("/:id", authMiddleware.CheckAuth, boardHandler.GetSingleBoard)
			boardRoutes.DELETE("/:id", authMiddleware.CheckAuth, boardHandler.DeleteBoard)
			boardRoutes.POST("/:id/:idU", authMiddleware.CheckAuth, boardHandler.AppendUserToBoard)
			boardRoutes.PUT("/:id/upload", authMiddleware.CheckAuth, boardHandler.SaveImage)
			boardRoutes.GET("/:id"+routes.ListRoute, authMiddleware.CheckAuth, listHandler.GetLists)
			boardRoutes.POST("/:id"+routes.ListRoute, authMiddleware.CheckAuth, listHandler.CreateList)
			boardRoutes.POST("/:id"+routes.ListRoute+"/:idL"+routes.TaskRoute, authMiddleware.CheckAuth, taskHandler.CreateTask)
			boardRoutes.GET("/append/:link", authMiddleware.CheckAuth, boardHandler.AppendUserToBoardByLink)
		}
		listRoutes := router.Group(routes.ListRoute)
		{
			listRoutes.GET("/:id", authMiddleware.CheckAuth, listHandler.GetSingleList)
			listRoutes.PUT("/:id", authMiddleware.CheckAuth, listHandler.RefactorList)
			listRoutes.DELETE("/:id", authMiddleware.CheckAuth, listHandler.DeleteList)
			listRoutes.GET("/:id"+routes.TaskRoute, authMiddleware.CheckAuth, taskHandler.GetTasks)
		}
		taskRoutes := router.Group(routes.TaskRoute)
		{
			taskRoutes.GET("", authMiddleware.CheckAuth, taskHandler.GetImportantTasks)
			taskRoutes.GET("/:id", authMiddleware.CheckAuth, taskHandler.GetSingleTask)
			taskRoutes.PUT("/:id", authMiddleware.CheckAuth, taskHandler.RefactorTask)
			taskRoutes.DELETE("/:id", authMiddleware.CheckAuth, taskHandler.DeleteTask)
			taskRoutes.POST("/:id/:idU", authMiddleware.CheckAuth, taskHandler.AppendUserToTask)
			taskRoutes.DELETE("/:id/:idU", authMiddleware.CheckAuth, taskHandler.DeleteUserFromTask)
			taskRoutes.GET("/:id"+routes.CheckListRoute, authMiddleware.CheckAuth, checkListHandler.GetCheckLists)
			taskRoutes.POST("/:id"+routes.CheckListRoute, authMiddleware.CheckAuth, checkListHandler.CreateCheckList)
			taskRoutes.GET("/:id"+routes.CommentRoute, authMiddleware.CheckAuth, commentHandler.GetComments)
			taskRoutes.POST("/:id"+routes.CommentRoute, authMiddleware.CheckAuth, commentHandler.CreateComment)
			taskRoutes.POST("/:id"+routes.AttachmentRoute, authMiddleware.CheckAuth, attachmentHandler.CreateAttachment)
			taskRoutes.GET("/append/:link", authMiddleware.CheckAuth, taskHandler.AppendUserToTaskByLink)
		}
		checkListRoutes := router.Group(routes.CheckListRoute)
		{
			checkListRoutes.GET("/:id", authMiddleware.CheckAuth, checkListHandler.GetSingleCheckList)
			checkListRoutes.PUT("/:id", authMiddleware.CheckAuth, checkListHandler.RefactorCheckList)
			checkListRoutes.DELETE("/:id", authMiddleware.CheckAuth, checkListHandler.DeleteCheckList)
			checkListRoutes.GET("/:id"+routes.CheckListItemRoute, authMiddleware.CheckAuth, checkListItemHandler.GetCheckListItems)
			checkListRoutes.POST("/:id"+routes.CheckListItemRoute, authMiddleware.CheckAuth, checkListItemHandler.CreateCheckListItem)
		}
		checkListItemRoutes := router.Group(routes.CheckListItemRoute)
		{
			checkListItemRoutes.GET("/:id", authMiddleware.CheckAuth, checkListItemHandler.GetSingleCheckListItem)
			checkListItemRoutes.PUT("/:id", authMiddleware.CheckAuth, checkListItemHandler.RefactorCheckListItem)
			checkListItemRoutes.DELETE("/:id", authMiddleware.CheckAuth, checkListItemHandler.DeleteCheckListItem)
		}
		commentRoutes := router.Group(routes.CommentRoute)
		{
			commentRoutes.GET("/:id", authMiddleware.CheckAuth, commentHandler.GetSingleComment)
			commentRoutes.PUT("/:id", authMiddleware.CheckAuth, commentHandler.RefactorComment)
			commentRoutes.DELETE("/:id", authMiddleware.CheckAuth, commentHandler.DeleteComment)
		}
		attachmentRoutes := router.Group(routes.AttachmentRoute)
		{
			attachmentRoutes.GET("/:id", authMiddleware.CheckAuth, attachmentHandler.GetSingleAttachment)
			attachmentRoutes.DELETE("/:id", authMiddleware.CheckAuth, attachmentHandler.DeleteAttachment)
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
