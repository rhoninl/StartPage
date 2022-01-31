package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"main/Controller"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("Views/*.html")
	router.StaticFS("src", http.Dir("Views/src"))
	router.Handle("GET", "/", Controller.GetMainPage)
	router.Run(":80")
}
