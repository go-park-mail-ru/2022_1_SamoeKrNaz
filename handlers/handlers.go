package handlers

import (
	"PLANEXA_backend/errors"
	boardModel "PLANEXA_backend/models/board_model"
	sessionModel "PLANEXA_backend/models/session_model"
	userModel "PLANEXA_backend/models/user_model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var cookieTime = 604800

func Login(c *gin.Context) {
	var user userModel.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}
	//token, err := c.Cookie("token")

	for _, userDB := range userModel.UserList {
		if userDB.Username == user.Username {
			if userDB.Password == user.Password {
				token := generateSessionToken()
				c.SetCookie("token", token, cookieTime, "", "", false, true)
				sessionModel.SessionList = append(sessionModel.SessionList, sessionModel.Session{UserId: userDB.Id, CookieValue: token})
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
	var user userModel.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}
	//token, err := c.Cookie("token")

	for _, userDB := range userModel.UserList {
		if userDB.Username == user.Username {
			c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUsernameExist), gin.H{"error": customErrors.ErrUsernameExist.Error()})
			return
		}
	}
	userModel.UserList = append(userModel.UserList, userModel.User{Id: userModel.UserID, Username: user.Username, Password: user.Password})
	token := generateSessionToken()
	c.SetCookie("token", token, cookieTime, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"is_registered": true})
	sessionModel.SessionList = append(sessionModel.SessionList, sessionModel.Session{UserId: userModel.UserID, CookieValue: token})
	userModel.UserID++
	return
}

func GetBoards(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	for _, sess := range sessionModel.SessionList {
		if token == sess.CookieValue {
			c.JSON(http.StatusOK, boardModel.BoardList)
			return
		}
	}
	c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
	return
}

func generateSessionToken() string {
	return uuid.NewString()
}
