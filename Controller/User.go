package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/utils"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var user utils.UserInfo
	c.BindJSON(&user)
	if user.UserName == "" || user.PassWord == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "Wrong Request"}})
		return
	}
	template := `Select UserId From startpage.User Where UserName = ? limit 1`
	result, err := DB().Query(template, user.UserName)
	if err != nil {
		log.Println("[Mysql] Query User make mistake")
		c.JSON(http.StatusInternalServerError, gin.H{"data": gin.H{"message": "DataBase make mistake"}})
		return
	}
	defer result.Close()
	if result.Next() {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "This Account Have Been Created"}})
		return
	}
	template = `Insert Into startpage.User Set UserName = ?,Password=?,CreatedTime=?`
	stmt, err := DB().Prepare(template)
	defer stmt.Close()
	rows, err := stmt.Exec(user.UserName, user.PassWord, time.Now())
	if err != nil {
		fmt.Println("[Mysql]Create User Default", err)
		DoLog(-1, c.ClientIP(), "Register Default")
		c.JSON(http.StatusInternalServerError, gin.H{"data": gin.H{"message": "DataBase make mistake"}})
		return
	}
	user.Id, _ = rows.LastInsertId()
	DoLog(user.Id, c.ClientIP(), "Register Success")
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": "Register Success!"}})
}

func Login(c *gin.Context) {
	var user utils.UserInfo
	currentPassword := ""
	c.BindJSON(&user)
	if user.UserName == "" || user.PassWord == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "Password Or Account is Empty"}})
		return
	}
	currentPassword, err := RDB().Get(user.UserName).Result()
	if err != nil {
		template := `Select Password From startpage.User Where UserName = ? limit 1`
		result, err := DB().Query(template, user.UserName)
		if err != nil {
			log.Println("[Mysql] Query User make mistake")
			c.JSON(http.StatusInternalServerError, gin.H{"data": gin.H{"message": "DataBase make mistake"}})
			return
		}
		defer result.Close()
		if !result.Next() {
			c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "Account Not Exists"}})
			return
		}
		result.Scan(&currentPassword)
	}
	if currentPassword == user.PassWord {
		RDB().Del(user.UserName)
		token, _ := utils.SetToken(user.Id)
		c.SetCookie("token", token, utils.MaxAge, "/", "/", true, true)
		c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": "Login Success"}})
		return
	}
	if err != nil {
		RDB().Set(user.UserName, currentPassword, time.Minute*3)
	}
	c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "Wrong Password"}})
}
