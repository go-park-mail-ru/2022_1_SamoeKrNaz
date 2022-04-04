package main

import (
	"PLANEXA_backend/handlers"
	"PLANEXA_backend/middleware"
	"PLANEXA_backend/models"
	planexa_redis "PLANEXA_backend/redis"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/routes"
	"PLANEXA_backend/usecases/impl"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=Planexa password=WEB21Planexa dbname=DB_Planexa port=5432",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Board{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.List{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initRedis() (client *planexa_redis.RedisConnect) {
	client = planexa_redis.RedisConnect{}.ConnectToRedis()
	return
}

func initRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://planexa.netlify.app", "http://89.208.199.114:3000", "http://89.208.199.114:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	db, _ := initDB()
	redis := initRedis()

	// создание хендлера
	userRepository := repositories.MakeUserRepository(db)
	userHandler := handlers.MakeUserHandler(impl.MakeUserUsecase(userRepository, redis))

	mainRoutes := router.Group(routes.HomeRoute)
	{
		boardRoutes := router.Group(routes.BoardRoute)
		{
			boardRoutes.POST("", middleware.CheckAuth, handlers.CreateBoard)
			boardRoutes.PUT("", middleware.CheckAuth, handlers.RefactorBoard)
			boardRoutes.GET("/:id", middleware.CheckAuth, handlers.GetSingleBoard)
			boardRoutes.DELETE("/:id", middleware.CheckAuth, handlers.DeleteBoard)
			boardRoutes.GET("/:id"+routes.ListRoute, middleware.CheckAuth, handlers.GetLists)
		}
		listRoutes := router.Group(routes.ListRoute)
		{
			listRoutes.GET("/:id", middleware.CheckAuth, handlers.GetSingleList)
			listRoutes.POST("", middleware.CheckAuth, handlers.CreateList)
			listRoutes.PUT("/:id", middleware.CheckAuth, handlers.RefactorList)
			listRoutes.DELETE("/:id", middleware.CheckAuth, handlers.DeleteList)
			listRoutes.GET("/:id"+routes.TaskRoute, middleware.CheckAuth, handlers.GetTasks)
		}
		taskRoutes := router.Group(routes.TaskRoute)
		{
			taskRoutes.GET("/:id", middleware.CheckAuth, handlers.GetSingleTask)
			taskRoutes.POST("", middleware.CheckAuth, handlers.CreateTask)
			taskRoutes.PUT("/:id", middleware.CheckAuth, handlers.RefactorTask)
			taskRoutes.DELETE("/:id", middleware.CheckAuth, handlers.DeleteTask)
		}
		mainRoutes.POST(routes.LoginRoute, handlers.Login)
		mainRoutes.GET("", middleware.CheckAuth, handlers.GetBoards)
		mainRoutes.POST(routes.RegisterRoute, handlers.Register)
		mainRoutes.DELETE(routes.LogoutRoute, handlers.Logout)
		mainRoutes.GET(routes.ProfileRoute+"/:id", middleware.CheckAuth, handlers.GetInfo)

	}
	return router
}
