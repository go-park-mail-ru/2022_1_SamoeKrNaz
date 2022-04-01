package main

import (
	"PLANEXA_backend/handlers"
	"PLANEXA_backend/middleware"
	"PLANEXA_backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://planexa.netlify.app", "http://89.208.199.114:3000", "http://89.208.199.114:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.LoginRoute, handlers.Login)
		mainRoutes.GET("", middleware.CheckAuth, handlers.GetBoards)
		mainRoutes.POST(routes.RegisterRoute, handlers.Register)
		mainRoutes.DELETE(routes.LogoutRoute, handlers.Logout)
		mainRoutes.POST(routes.BoardRoute, middleware.CheckAuth, handlers.CreateBoard)
		mainRoutes.PUT(routes.BoardRoute, middleware.CheckAuth, handlers.RefactorBoard)
		mainRoutes.GET(routes.BoardRoute+"/:id", middleware.CheckAuth, handlers.GetSingleBoard)
		mainRoutes.DELETE(routes.BoardRoute+"/:id", middleware.CheckAuth, handlers.DeleteBoard)
		mainRoutes.GET(routes.ProfileRoute+"/:id", middleware.CheckAuth, handlers.GetInfo)
		mainRoutes.GET(routes.BoardRoute+"/:id"+routes.ListRoute, middleware.CheckAuth, handlers.GetLists)
		mainRoutes.GET(routes.ListRoute+"/:id", middleware.CheckAuth, handlers.GetSingleList)
		mainRoutes.POST(routes.ListRoute, middleware.CheckAuth, handlers.CreateList)
		mainRoutes.PUT(routes.ListRoute+"/:id", middleware.CheckAuth, handlers.RefactorList)
		mainRoutes.DELETE(routes.ListRoute+"/:id", middleware.CheckAuth, handlers.DeleteList)
		mainRoutes.GET(routes.ListRoute+"/:id"+routes.TaskRoute, middleware.CheckAuth, handlers.GetTasks)
		mainRoutes.GET(routes.TaskRoute+"/:id", middleware.CheckAuth, handlers.GetSingleTask)
		mainRoutes.POST(routes.TaskRoute, middleware.CheckAuth, handlers.CreateTask)
		mainRoutes.PUT(routes.TaskRoute+"/:id", middleware.CheckAuth, handlers.RefactorTask)
		mainRoutes.DELETE(routes.TaskRoute+"/:id", middleware.CheckAuth, handlers.DeleteTask)

	}
	return router
}
