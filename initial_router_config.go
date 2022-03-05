package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/handlers"
	"main/routes"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	mainRoutes := router.Group("/api")
	{
		mainRoutes.POST(routes.LoginRoute, handlers.Login)
		mainRoutes.GET(routes.HomeRoute, handlers.GetBoards)
		mainRoutes.POST(routes.RegisterRoute, handlers.Register)
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://planexa.netlify.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true

	router.Use(cors.New(config))
	return router
}
