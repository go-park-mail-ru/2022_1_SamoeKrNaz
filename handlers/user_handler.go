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
	threeDays = 72 * time.Hour
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

	userId, token, err := userHandler.usecase.Login(user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	user, err = userHandler.usecase.GetInfoById(userId)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	expiration := time.Now().Add(threeDays)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, &user)
	return
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	userId, token, err := userHandler.usecase.Register(user)

	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	user, err = userHandler.usecase.GetInfoById(userId)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	expiration := time.Now().Add(threeDays)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusCreated, &user)
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

func (userHandler *UserHandler) GetInfoById(c *gin.Context) {
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

	user, err := userHandler.usecase.GetInfoById(uint(userId))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func (userHandler *UserHandler) GetInfoByCookie(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	user, err := userHandler.usecase.GetInfoById(uint(userId.(uint64)))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func (userHandler *UserHandler) SaveAvatar(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	header, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	user := new(models.User)
	user.IdU = userId.(uint)
	user.ImgAvatar = header.Filename

	err = userHandler.usecase.SaveAvatar(user, header)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"avatar_path": user.ImgAvatar})
	return
}

func (userHandler *UserHandler) RefactorProfile(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	if uint(userId.(uint64)) != user.IdU {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	err = userHandler.usecase.RefactorProfile(user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"avatar_path": user.ImgAvatar})
	return
}
