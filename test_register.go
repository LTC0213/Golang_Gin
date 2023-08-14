package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string   `form:"username"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
	Gender   string   `form:"gender"`
	City     string   `form:"city"`
}

func Regsiter(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	c.String(200, "User:%s", user)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.POST("/register", Regsiter)
	e.GET("/register", GoRegister)
	e.Run()
}
