package Middlewares

import (
	"github.com/gin-gonic/gin"
	"main/utils"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusPermanentRedirect, "/login")
		}
		userid, err := utils.CheckToken(token)
		if err == nil {
			c.Set("userId", userid)
		}
		//token, _ := utils.SetToken( )
		//c.SetCookie("token", token, utils.MaxAge, "/", "/", true, true)
	}
}
