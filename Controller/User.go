package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RegForm struct {
	UserName string `json:UserName`
	PassWord string `json:Password`
}

func Register(c *gin.Context) {
	var user RegForm
	c.BindJSON(&user)
	if user.UserName == "" || user.PassWord == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "Wrong Request"}})
		return
	}
	template := `Insert Into Table User Set UserName = ?,Password=?,CreatedTime=?`
	_, err := DB().Query(template, user.UserName, user.PassWord, time.Now())
	if err != nil {
		fmt.Println("[Mysql]Create User Default", err)
		DoLog(-1, c.ClientIP(), "Register Default")
		c.JSON(http.StatusInternalServerError, gin.H{"data": gin.H{"message": "DataBase make mistake"}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": "Register Success!"}})
}
