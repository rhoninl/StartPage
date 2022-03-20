package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
