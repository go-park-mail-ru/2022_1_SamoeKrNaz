package middleware

import (
	"PLANEXA_backend/auth_microservice/server/handler"
	"context"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	redisRep handler.AuthCheckerClient
	ctx      context.Context
}

func CreateMiddleware(rep handler.AuthCheckerClient) *Middleware {
	return &Middleware{redisRep: rep, ctx: context.Background()}
}

func (mw *Middleware) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД
	userId, err := mw.redisRep.Get(mw.ctx, &handler.SessionValue{Value: token})
	if err != nil {
		return
	}
	c.Set("Auth", userId)
}
