package main

import (
	"PLANEXA_backend/handlers"
	"PLANEXA_backend/middleware"
	"PLANEXA_backend/models"
	impl_rep "PLANEXA_backend/repositories/impl"
	"PLANEXA_backend/routes"
	"PLANEXA_backend/usecases/impl"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	err = db.AutoMigrate(&models.User{}, &models.Board{}, &models.List{}, &models.Task{})
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
	redis := impl_rep.ConnectToRedis()

	// создание репозиториев
	userRepository := impl_rep.MakeUserRepository(db)
	taskRepository := impl_rep.MakeTaskRepository(db)
	listRepository := impl_rep.MakeListRepository(db)
	boardRepository := impl_rep.MakeBoardRepository(db)

	authMiddleware := middleware.CreateMiddleware(redis)

	userHandler := handlers.MakeUserHandler(impl.MakeUserUsecase(userRepository, redis))
	taskHandler := handlers.MakeTaskHandler(impl.MakeTaskUsecase(taskRepository, boardRepository, listRepository))
	boardHandler := handlers.MakeBoardHandler(impl.MakeBoardUsecase(boardRepository, listRepository))
	listHandler := handlers.MakeListHandler(impl.MakeListUsecase(listRepository, boardRepository))

	mainRoutes := router.Group(routes.HomeRoute)
	{
		boardRoutes := router.Group(routes.BoardRoute)
		{
			boardRoutes.POST("", authMiddleware.CheckAuth, boardHandler.CreateBoard)
			boardRoutes.PUT("/:id", authMiddleware.CheckAuth, boardHandler.RefactorBoard)
			boardRoutes.GET("/:id", authMiddleware.CheckAuth, boardHandler.GetSingleBoard)
			boardRoutes.DELETE("/:id", authMiddleware.CheckAuth, boardHandler.DeleteBoard)
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
			taskRoutes.GET("/:id", authMiddleware.CheckAuth, taskHandler.GetSingleTask)
			taskRoutes.PUT("/:id", authMiddleware.CheckAuth, taskHandler.RefactorTask)
			taskRoutes.DELETE("/:id", authMiddleware.CheckAuth, taskHandler.DeleteTask)
		}
		mainRoutes.POST(routes.LoginRoute, userHandler.Login)
		mainRoutes.GET("/get"+routes.BoardRoute+"s", authMiddleware.CheckAuth, boardHandler.GetBoards)
		mainRoutes.POST(routes.RegisterRoute, userHandler.Register)
		mainRoutes.DELETE(routes.LogoutRoute, userHandler.Logout)
		mainRoutes.GET(routes.ProfileRoute+"/:id", authMiddleware.CheckAuth, userHandler.GetInfoById)
		mainRoutes.GET(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.GetInfoByCookie)
		mainRoutes.PUT(routes.ProfileRoute+"/upload", authMiddleware.CheckAuth, userHandler.SaveAvatar)
		mainRoutes.PUT(routes.ProfileRoute, authMiddleware.CheckAuth, userHandler.RefactorProfile)
	}
	return router, nil
}
