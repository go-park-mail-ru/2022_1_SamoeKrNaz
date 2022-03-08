package main

import (
	"PLANEXA_backend/handlers"
	"PLANEXA_backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://planexa.netlify.app", "https://89.208.199.114:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	mainRoutes := router.Group(routes.HomeRoute)
	{
		mainRoutes.POST(routes.LoginRoute, handlers.Login)
		mainRoutes.GET("", handlers.GetBoards)
		mainRoutes.POST(routes.RegisterRoute, handlers.Register)
		mainRoutes.DELETE(routes.LogoutRoute, handlers.Logout)
	}
	return router
}
