package middleware

import (
	"PLANEXA_backend/repositories"
	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД
	redis := repositories.ConnectToRedis()
	userId, err := redis.GetSession(token)
	if err != nil {
		return
	}
	c.Set("Auth", userId)
	return
}
