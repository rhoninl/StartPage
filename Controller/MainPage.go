package Controller

import (
	"github.com/gin-gonic/gin"
	"main/utils"
	"net/http"
	"strconv"
)

func MainPage(c *gin.Context) {
	userId, exists := c.Get("userId")
	background := ""
	if !exists {
		background = "./src/background.jpg?v=1.1"
	} else {
		background, _ = utils.GetUserBg(userId.(int64))
		if background == "" {
			background = "./src/background.jpg?v=1.1"
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"isLogin":    strconv.FormatBool(exists),
		"background": background,
	})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", gin.H{})
}

func LogOut(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	_, err = utils.CheckToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	token, err = utils.DeleteToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.SetCookie("token", token, -1, "/", "/", true, true)
	c.JSON(http.StatusOK, gin.H{})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Register.html", gin.H{})
}

func Notice(c *gin.Context) {
	c.HTML(http.StatusOK, "notice.html", gin.H{})
}

func Setting(c *gin.Context) {
	c.HTML(http.StatusOK, "Setting.html", gin.H{})
}
