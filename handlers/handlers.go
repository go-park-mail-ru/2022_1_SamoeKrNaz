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
var cookieTime = 604800
var lock = sync.RWMutex{}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	//lock.RLock()
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
	//lock.RUnlock()

	c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
	return
}

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}
	err = utils.CheckPassword(user.Password)
	if err != nil {
		//lock.RLock()
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrPassword), gin.H{"error": err.Error()})
		//lock.RUnlock()
	}

	//lock.RLock()
	for _, userDB := range models.UserList {
		if userDB.Username == user.Username {
			c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUsernameExist), gin.H{"error": customErrors.ErrUsernameExist.Error()})
			return
		}
	}
	models.UserList = append(models.UserList, models.User{Id: models.UserID, Username: user.Username, Password: user.Password})
	//lock.RUnlock()
	token := generateSessionToken()
	c.SetCookie("token", token, cookieTime, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"is_registered": true})
	//lock.RLock()
	models.SessionList = append(models.SessionList, models.Session{UserId: models.UserID, CookieValue: token})
	models.UserID++
	//lock.RUnlock()
	return
}

func GetBoards(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	//lock.RLock()
	for _, sess := range models.SessionList {
		if token == sess.CookieValue {
			c.JSON(http.StatusOK, models.BoardList)
			return
		}
	}
	//lock.RUnlock()
	c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"ERROR": customErrors.ErrUnauthorized.Error()})
	return
}

func generateSessionToken() string {
	return uuid.NewString()
}
