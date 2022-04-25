package middleware

import (
	"PLANEXA_backend/repositories"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	redisRep repositories.RedisRepository
}

func CreateMiddleware(rep repositories.RedisRepository) *Middleware {
	return &Middleware{redisRep: rep}
}

func (mw *Middleware) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД
	userId, err := mw.redisRep.GetSession(token)
	if err != nil {
		return
	}
	c.Set("Auth", userId)
}
