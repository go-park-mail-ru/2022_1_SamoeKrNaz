package handlers

import (
	"PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

// 3 days
var (
	cookieTime = 604800
	lockUser   = sync.RWMutex{}
	lockSess   = sync.RWMutex{}
)

func Login(c *gin.Context) {
	lockUser.RLock()
	lockSess.Lock()
	defer lockUser.RUnlock()
	defer lockSess.Unlock()

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	for _, userDB := range models.UserList {
		if userDB.Username == user.Username {
			if userDB.Password == user.Password {
				token := generateSessionToken()
				c.SetCookie("token", token, cookieTime, "", "", false, true)
				models.SessionList = append(models.SessionList, models.Session{UserId: userDB.Id, CookieValue: token})
				c.JSON(http.StatusOK, gin.H{"is_logged": true})
				return
			} else {
				break
			}
		}
	}

	c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
	return
}

func Register(c *gin.Context) {
	lockUser.Lock()
	lockSess.Lock()
	defer lockUser.Unlock()
	defer lockSess.Unlock()

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}
	err = utils.CheckPassword(user.Password)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}

	for _, userDB := range models.UserList {
		if userDB.Username == user.Username {
			c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUsernameExist), gin.H{"error": customErrors.ErrUsernameExist.Error()})
			return
		}
	}
	models.UserList = append(models.UserList, models.User{Id: models.UserID, Username: user.Username, Password: user.Password})
	token := generateSessionToken()
	c.SetCookie("token", token, cookieTime, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"is_registered": true})
	models.SessionList = append(models.SessionList, models.Session{UserId: models.UserID, CookieValue: token})
	models.UserID++
	return
}

func GetBoards(c *gin.Context) {
	lockSess.Lock()
	defer lockSess.Unlock()

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	for _, sess := range models.SessionList {
		if token == sess.CookieValue {
			c.JSON(http.StatusOK, models.BoardList)
			return
		}
	}
	c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"ERROR": customErrors.ErrUnauthorized.Error()})
	return
}

func Logout(c *gin.Context) {
	lockSess.Lock()
	defer lockSess.Unlock()
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
	}

	for i, sess := range models.SessionList {
		if sess.CookieValue == token {
			models.SessionList[i] = models.SessionList[len(models.SessionList)-1]
			models.SessionList = models.SessionList[:len(models.SessionList)-1]
			c.JSON(http.StatusOK, gin.H{"Is_okay": true})
			return
		}
	}
	c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"err": customErrors.ErrUnauthorized.Error()})
}

func generateSessionToken() string {
	return uuid.NewString()
}
