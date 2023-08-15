package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Checked  string `form:"checked"`
}

func Login(c *gin.Context) {
	c.HTML(200, "static.html", nil)
}

func DoLogin(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	c.String(200, "User:%s", user)
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.Static("/assets", "./assets")
	//e.StaticFS("/croot", http.Dir("c:/"))
	//e.StaticFile("/favicon.ico", "./assets/favicon.ico")

	e.GET("/login", Login)
	e.POST("/login", DoLogin)
	e.Run()
}
