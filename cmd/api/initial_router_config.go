package main

import (
	"PLANEXA_backend/auth_microservice/server/handler"
	customErrors "PLANEXA_backend/errors"
	handlers2 "PLANEXA_backend/main_microservice/handlers"
	"PLANEXA_backend/main_microservice/middleware"
	impl3 "PLANEXA_backend/main_microservice/repositories/impl"
	impl2 "PLANEXA_backend/main_microservice/usecases/impl"
	"PLANEXA_backend/models"
	"PLANEXA_backend/routes"
	handler_user "PLANEXA_backend/user_microservice/server_user/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("host=postgres user=Planexa password=WEB21Planexa dbname=DB_Planexa port=5432"))
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
	gin.DisableConsoleColor()
	f, err := os.Create("gin.log")
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

	db, err := initDB()
	if err != nil {
		return nil, err
	}
	grpcConn, err := grpc.Dial(
		"2022_1_samoekrnaz_session_1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, customErrors.ErrNoAccess
	}

	sessService := impl3.CreateRepo(handler.NewAuthCheckerClient(grpcConn))

	if err != nil {
		return nil, err
	}
	grpcConnUser, err := grpc.Dial(
		"2022_1_samoekrnaz_user_microservice_1:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, customErrors.ErrNoAccess
	}
	userService := impl3.MakeUserRepository(db, handler_user.NewUserServiceClient(grpcConnUser))

	monitor := ginmetrics.GetMonitor()
	if err != nil {
		return nil, err
	}
	monitor.SetMetricPath("/metrics")
	monitor.SetSlowTime(10)
	monitor.Use(router)

	// создание репозиториев
	taskRepository := impl3.MakeTaskRepository(db)
	listRepository := impl3.MakeListRepository(db)
	boardRepository := impl3.MakeBoardRepository(db)
	checkListRepository := impl3.MakeCheckListRepository(db)
	checkListItemRepository := impl3.MakeCheckListItemRepository(db)
	commentRepository := impl3.MakeCommentRepository(db)

	authMiddleware := middleware.CreateMiddleware(sessService)

	userHandler := handlers2.MakeUserHandler(impl2.MakeUserUsecase(userService, sessService))
	taskHandler := handlers2.MakeTaskHandler(impl2.MakeTaskUsecase(taskRepository, boardRepository, listRepository, userService, checkListRepository, commentRepository))
	boardHandler := handlers2.MakeBoardHandler(impl2.MakeBoardUsecase(boardRepository, listRepository, taskRepository, checkListRepository, userService, commentRepository))
	listHandler := handlers2.MakeListHandler(impl2.MakeListUsecase(listRepository, boardRepository))
	checkListHandler := handlers2.MakeCheckListHandler(impl2.MakeCheckListUsecase(checkListRepository, taskRepository))
	checkListItemHandler := handlers2.MakeCheckListItemHandler(impl2.MakeCheckListItemUsecase(checkListItemRepository, checkListRepository, taskRepository))
	commentHandler := handlers2.MakeCommentHandler(impl2.MakeCommentUsecase(commentRepository, taskRepository, userService))
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
