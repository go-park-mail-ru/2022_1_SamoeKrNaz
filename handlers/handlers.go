package handlers

import (
	"PLANEXA_backend/errors"
	boardModel "PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// 3 days
var cookieTime = 604800

func Login(c *gin.Context) {
	var user boardModel.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	for _, userDB := range boardModel.UserList {
		if userDB.Username == user.Username {
			if userDB.Password == user.Password {
				token := generateSessionToken()
				c.SetCookie("token", token, cookieTime, "", "", false, true)
				boardModel.SessionList = append(boardModel.SessionList, boardModel.Session{UserId: userDB.Id, CookieValue: token})
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
	var user boardModel.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	for _, userDB := range boardModel.UserList {
		if userDB.Username == user.Username {
			c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUsernameExist), gin.H{"error": customErrors.ErrUsernameExist.Error()})
			return
		}
	}
	boardModel.UserList = append(boardModel.UserList, boardModel.User{Id: boardModel.UserID, Username: user.Username, Password: user.Password})
	token := generateSessionToken()
	c.SetCookie("token", token, cookieTime, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"is_registered": true})
	boardModel.SessionList = append(boardModel.SessionList, boardModel.Session{UserId: boardModel.UserID, CookieValue: token})
	boardModel.UserID++
	return
}

func GetBoards(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	for _, sess := range boardModel.SessionList {
		if token == sess.CookieValue {
			c.JSON(http.StatusOK, boardModel.BoardList)
			return
		}
	}
	c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"ERROR": customErrors.ErrUnauthorized.Error()})
	return
}

func generateSessionToken() string {
	return uuid.NewString()
}
