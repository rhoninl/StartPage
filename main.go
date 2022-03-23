package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"main/Controller"
	"main/Middlewares"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("Views/*.html")
	router.StaticFile("/favicon.ico", "./Views/src/favicon.ico")
	router.StaticFS("src", http.Dir("Views/src"))
	router.Handle("GET", "/", Middlewares.Auth(), Controller.MainPage)
	router.Handle("GET", "/Login", Controller.LoginPage)
	router.Handle("GET", "/Register", Controller.RegisterPage)
	router.Handle("POST", "/Login", Controller.Login)
	router.Handle("POST", "/Register", Controller.Register)
	router.Handle("POST", "/LogOut", Controller.LogOut)
	router.Run(":8081")
}
