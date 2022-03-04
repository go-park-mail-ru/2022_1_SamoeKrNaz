package handlers

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"math/rand"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_logged": false})
		return
	}
	//token, err := c.Cookie("token")

	for _, userDB := range models.UserList {
		if userDB.Username == user.Username {
			if userDB.Password == user.Password {
				token := generateSessionToken()
				c.SetCookie("token", token, 3600, "", "", false, true)
				models.SessionList = append(models.SessionList, models.Session{SessionId: userDB.Id, CookieId: token})
				c.JSON(http.StatusOK, gin.H{"is_logged": true})
				return
			}
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"is_logged": false})
	return
}

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_registered": false})
		return
	}
	//token, err := c.Cookie("token")

	for _, userDB := range models.UserList {
		if userDB.Username == user.Username {
			c.JSON(http.StatusConflict, gin.H{"is_registered": false})
			return
		}
	}
	models.UserList = append(models.UserList, models.User{Id: models.UserID, Username: user.Username, Password: user.Password})
	token := generateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"is_registered": true})
	models.SessionList = append(models.SessionList, models.Session{SessionId: models.UserID, CookieId: token})
	models.UserID++
	return
}

func GetBoards(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"isOkay": false})
		return
	}

	for _, sess := range models.SessionList {
		if token == sess.CookieId {
			c.JSON(http.StatusOK, models.BoardList)
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"isOkay": false})
	return
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
