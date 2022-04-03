package Middlewares

import (
	"github.com/gin-gonic/gin"
	"main/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err == nil {
			// 获取userId
			userid, err := utils.CheckToken(token)
			if err == nil {
				//获取成功则向后传输userId
				c.Set("userId", userid)
			}
		}
		//token, _ := utils.SetToken( )
		//c.SetCookie("token", token, utils.MaxAge, "/", "/", true, true)
	}
}
