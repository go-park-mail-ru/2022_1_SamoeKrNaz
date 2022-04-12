package middleware

import (
	"PLANEXA_backend/repositories/impl"
	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД
	redis := impl.ConnectToRedis()
	userId, err := redis.GetSession(token)
	if err != nil {
		return
	}
	c.Set("Auth", userId)
	return
}
