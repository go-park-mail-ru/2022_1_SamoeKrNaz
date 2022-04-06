package main

import (
	"PLANEXA_backend/handlers"
	"PLANEXA_backend/middleware"
	"PLANEXA_backend/models"
	"PLANEXA_backend/redis"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/routes"
	"PLANEXA_backend/usecases/impl"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=Planexa password=WEB21Planexa dbname=DB_Planexa port=5432"))
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{}, &models.Board{}, &models.List{}, &models.Task{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://planexa.netlify.app", "http://89.208.199.114:3000", "http://89.208.199.114:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	db, err := initDB()
	if err != nil {
		return nil
	}
	redis := planexa_redis.ConnectToRedis()

	// создание репозиториев
	userRepository := repositories.MakeUserRepository(db)
	taskRepository := repositories.MakeTaskRepository(db)
	listRepository := repositories.MakeListRepository(db)
	boardRepository := repositories.MakeBoardRepository(db)

	userHandler := handlers.MakeUserHandler(impl.MakeUserUsecase(userRepository, redis))
	taskHandler := handlers.MakeTaskHandler(impl.MakeTaskUsecase(taskRepository))
	boardHandler := handlers.MakeBoardHandler(impl.MakeBoardUsecase(boardRepository))
	listHandler := handlers.MakeListHandler(impl.MakeListUsecase(listRepository))

	mainRoutes := router.Group(routes.HomeRoute)
	{
		boardRoutes := router.Group(routes.BoardRoute)
		{
			boardRoutes.POST("", middleware.CheckAuth, boardHandler.CreateBoard)
			boardRoutes.PUT("", middleware.CheckAuth, boardHandler.RefactorBoard)
			boardRoutes.GET("/:id", middleware.CheckAuth, boardHandler.GetSingleBoard)
			boardRoutes.DELETE("/:id", middleware.CheckAuth, boardHandler.DeleteBoard)
			boardRoutes.GET("/:id"+routes.ListRoute, middleware.CheckAuth, listHandler.GetLists)
			boardRoutes.POST("/:id"+routes.ListRoute, middleware.CheckAuth, listHandler.CreateList)
			boardRoutes.POST("/:idB"+routes.ListRoute+"/:idL"+routes.TaskRoute, middleware.CheckAuth, taskHandler.CreateTask)

		}
		listRoutes := router.Group(routes.ListRoute)
		{
			listRoutes.GET("/:id", middleware.CheckAuth, listHandler.GetSingleList)
			listRoutes.PUT("/:id", middleware.CheckAuth, listHandler.RefactorList)
			listRoutes.DELETE("/:id", middleware.CheckAuth, listHandler.DeleteList)
			listRoutes.GET("/:id"+routes.TaskRoute, middleware.CheckAuth, taskHandler.GetTasks)
		}
		taskRoutes := router.Group(routes.TaskRoute)
		{
			taskRoutes.GET("/:id", middleware.CheckAuth, taskHandler.GetSingleTask)
			taskRoutes.PUT("/:id", middleware.CheckAuth, taskHandler.RefactorTask)
			taskRoutes.DELETE("/:id", middleware.CheckAuth, taskHandler.DeleteTask)
		}
		mainRoutes.POST(routes.LoginRoute, userHandler.Login)
		mainRoutes.GET("", middleware.CheckAuth, boardHandler.GetBoards)
		mainRoutes.POST(routes.RegisterRoute, userHandler.Register)
		mainRoutes.DELETE(routes.LogoutRoute, userHandler.Logout)
		mainRoutes.GET(routes.ProfileRoute+"/:id", middleware.CheckAuth, userHandler.GetInfo)

	}
	return router
}
