package middleware

import (
	"PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	// Получаю сессии из БД

	for _, sess := range models.SessionList {
		if token == sess.CookieValue {
			//c.JSON(http.StatusOK, models.TasksAndBoards)
			c.Set("Auth", sess.UserId)
			return
		}
	}
	return
}

func CheckContentType(c *gin.Context) {
	content, check := c.Get("Content-Type")
	if !check {
		return
	} else {
		c.Set("content", content)
	}
}
