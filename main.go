package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	SessionId int    `json:"session_Id"`
	CookieId  string `json:"cookie_id"`
}

type Board struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Date        string `json:"date"`
}

var router *gin.Engine

var userID int
var userList []User
var sessionList []Session
var boardList []Board

func main() {

	router = gin.Default()

	InitializeRoutes()

	router.Run()
}

func InitializeRoutes() {
	router.POST("/api/login", login)
	router.GET("/api/", register)
	router.POST("/api/register", getBoards)
}

func login(c *gin.Context) {
	/*var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_logged": false})
		return
	}
	//token, err := c.Cookie("token")

	for _, userDB := range userList {
		if userDB.Username == user.Username {
			if userDB.Password == user.Password {
				token := generateSessionToken()
				c.SetCookie("token", token, 3600, "", "", false, true)
				sessionList = append(sessionList, Session{userDB.Id, token})
				c.JSON(http.StatusOK, gin.H{"is_logged": true})
				return
			}
		}
	}*/

	c.JSON(http.StatusUnauthorized, gin.H{"is_logged": false})
	return
}

func register(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_registered": false})
		return
	}
	//token, err := c.Cookie("token")

	for _, userDB := range userList {
		if userDB.Username == user.Username {
			c.JSON(http.StatusConflict, gin.H{"is_registered": false})
		}
	}
	userList = append(userList, User{userID, user.Username, user.Password})
	token := generateSessionToken()
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"is_registered": true})
	sessionList = append(sessionList, Session{userID, token})
	userID++
	return
}

func getBoards(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isOkay": false})
	}

	for _, sess := range sessionList {
		if token == sess.CookieId {
			c.JSON(http.StatusOK, boardList)
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"isOkay": false})
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
