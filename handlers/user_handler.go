package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var (
	CookieTime = 604800 // 3 days
)

type UserHandler struct {
	usecase usecases.UserUseCase
}

func MakeUserHandler(usecase usecases.UserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	// вызываю юзкейс

	token, err := userHandler.usecase.Login(user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	
	expiration := time.Now().Add(72 * time.Hour)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)
	return
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	token, err := userHandler.usecase.Register(user)

	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	expiration := time.Now().Add(72 * time.Hour)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)
	return
}

func (userHandler *UserHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	//usecase

	err = userHandler.usecase.Logout(token)

	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"err": err.Error()})
		return
	}

	c.SetCookie("token", token, -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Is_okay": true})
	return
}

func (userHandler *UserHandler) GetInfo(c *gin.Context) {
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

	user, err := userHandler.usecase.GetInfo(uint(userId))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
