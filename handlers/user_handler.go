package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

var (
	CookieTime = 604800 // 3 days
)

func generateSessionToken() string {
	return uuid.NewString()
}

type UserHandler struct {
	usecase *usecases.UserUsecase
}

func MakeUserHandler(usecase *usecases.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (*UserHandler) Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	// вызываю юзкейс

	token, err := usecases.Login(user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("token", token, CookieTime, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"is_logged": true})
	return
}

func (*UserHandler) Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	//usecase

	token, err := usecases.Register(user)

	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, CookieTime, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"is_registered": true})
	return
}

func (*UserHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	//usecase

	err = usecases.Logout(token)

	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"err": err.Error()})
		return
	}

	c.SetCookie("token", token, -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Is_okay": true})
	return
}

func (*UserHandler) GetInfo(c *gin.Context) {
	_, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}
	userId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	user, err := usecases.GetInfo(uint(userId))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
