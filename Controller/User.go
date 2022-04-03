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
	result, err := utils.DB().Query(template, user.UserName)
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
	stmt, err := utils.DB().Prepare(template)
	defer stmt.Close()
	rows, err := stmt.Exec(user.UserName, user.PassWord, time.Now())
	if err != nil {
		fmt.Println("[Mysql]Create User Default", err)
		utils.DoLog(-1, c.ClientIP(), "Register Default")
		c.JSON(http.StatusInternalServerError, gin.H{"data": gin.H{"message": "DataBase make mistake"}})
		return
	}
	user.Id, _ = rows.LastInsertId()
	utils.InitSetting(user.Id)
	utils.DoLog(user.Id, c.ClientIP(), "Register Success")
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
	currentPassword, err := utils.RDB().Get(user.UserName).Result()
	if err != nil {
		template := `Select Password From startpage.User Where UserName = ? limit 1`
		result, err := utils.DB().Query(template, user.UserName)
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
		utils.RDB().Del(user.UserName)
		userId, err := utils.GetUserIdByUserName(user.UserName)
		if err != nil {
			userId = -1
		}
		token, _ := utils.SetToken(userId)
		c.SetCookie("token", token, utils.MaxAge, "/", "/", false, true)
		c.JSON(http.StatusOK, gin.H{"data": gin.H{"message": "Login Success"}})
		return
	}
	if err != nil {
		utils.RDB().Set(user.UserName, currentPassword, time.Minute*3)
	}
	c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "Wrong Password"}})
}

func Favourite(c *gin.Context) {
	userId, exists := c.Get("userId")
	if exists {
		fmt.Println(userId)
	}
}

func SetSetting(c *gin.Context) {
	var userInfo utils.UserSetting
	c.Bind(&userInfo)
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "请求错误"}})
		return
	}
	userInfo.Id = userId.(int64)
	ok := utils.UpdateSetting(userInfo)
	if ok {
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"data": gin.H{"message": "设置失败"}})
}
