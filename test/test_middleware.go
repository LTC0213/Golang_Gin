package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMW(c *gin.Context) {
	c.String(200, "hello,%s", "ghz")
}

func MyMiddleware1(c *gin.Context) {
	fmt.Println("我的第一个中间件")
}

func MyMiddleware2(c *gin.Context) {
	fmt.Println("我的第二个中间件")
}

func main() {

	/* 	func Default() *Engine {
		debugPrintWARNINGDefault()
		engine := New()
		engine.Use(Logger(), Recovery())
		return engine
	} */
	// e := gin.Default()
	// e := gin.New()

	e := gin.Default()

	e.Use(MyMiddleware1, MyMiddleware2)

	e.GET("testmw", TestMW)

	e.Run()

}
