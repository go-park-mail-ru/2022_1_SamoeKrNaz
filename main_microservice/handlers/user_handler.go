package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
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
	err := easyjson.UnmarshalFromReader(c.Request.Body, &user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	// вызываю юзкейс

	userId, token, err := userHandler.usecase.Login(user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	user, err = userHandler.usecase.GetInfoById(userId)
	if err != nil {
		_ = c.Error(err)
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

	userJson, err := user.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	http.SetCookie(c.Writer, &cookie)
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	var user models.User
	err := easyjson.UnmarshalFromReader(c.Request.Body, &user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	userId, token, err := userHandler.usecase.Register(user)

	if err != nil {
		_ = c.Error(err)
		return
	}

	user, err = userHandler.usecase.GetInfoById(userId)
	if err != nil {
		_ = c.Error(err)
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

	userJson, err := user.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	http.SetCookie(c.Writer, &cookie)
	c.Data(http.StatusCreated, "application/json; charset=utf-8", userJson)
}

func (userHandler *UserHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	//usecase

	err = userHandler.usecase.Logout(token)

	if err != nil {
		_ = c.Error(err)
		return
	}

	var isOkay models.Is_okayIn
	isOkay.Is_okayInfo = true
	isOkayJson, err := isOkay.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.SetCookie("token", token, -1, "", "", false, true)
	c.Data(http.StatusOK, "application/json; charset=utf-8", isOkayJson)
}

func (userHandler *UserHandler) GetInfoById(c *gin.Context) {
	_, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}
	userId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	user, err := userHandler.usecase.GetInfoById(uint(userId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userJson, err := user.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (userHandler *UserHandler) GetInfoByCookie(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	user, err := userHandler.usecase.GetInfoById(uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userJson, err := user.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (userHandler *UserHandler) SaveAvatar(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	header, err := c.FormFile("avatar")
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	user := new(models.User)
	user.IdU = uint(userId.(uint64))
	user.ImgAvatar = header.Filename

	path, err := userHandler.usecase.SaveAvatar(user, header)

	if err != nil {
		_ = c.Error(err)
		return
	}

	var avatarPath models.Avatar
	avatarPath.AvatarPath = path
	avatarPathJson, err := avatarPath.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", avatarPathJson)

}

func (userHandler *UserHandler) RefactorProfile(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var user models.User
	err := easyjson.UnmarshalFromReader(c.Request.Body, &user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	user.IdU = uint(userId.(uint64))
	err = userHandler.usecase.RefactorProfile(user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var isUpdated models.Updated
	isUpdated.UpdatedInfo = true
	isUpdatedJson, err := isUpdated.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", isUpdatedJson)
}

func (userHandler *UserHandler) GetUsersLike(c *gin.Context) {
	_, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var user models.User
	err := easyjson.UnmarshalFromReader(c.Request.Body, &user)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	usersLike, err := userHandler.usecase.GetUsersLike(user.Username)
	if err != nil {
		_ = c.Error(err)
		return
	}

	newUsersLike := new(models.Users)
	*newUsersLike = *usersLike

	usersJson, err := newUsersLike.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", usersJson)
}
