package main

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	ctx.String(200, "hello %s", "LTC")
}

func main() {
	e := gin.Default()
	e.GET("/hello", Hello)
	e.Run(":8080")
}
