package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("Views/*.html")
	router.StaticFS("src", http.Dir("Views/src"))
	router.Handle("GET", "/", GetMainPage)
	router.Run(":80")
}
func GetMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
