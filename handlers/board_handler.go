package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBoards(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	//Получаю доски от БД
	boards, err := usecases.GetBoards(userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, boards)
	return
}
