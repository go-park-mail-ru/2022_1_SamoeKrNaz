package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/handlers"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	InitializeRoutes()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://planexa.netlify.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true

	router.Use(cors.New(config))
	cors.DefaultConfig()
	router.Run()
}

func InitializeRoutes() {
	router.POST("/api/login", handlers.Login)
	router.GET("/api/", handlers.GetBoards)
	router.POST("/api/register", handlers.Register)
}
