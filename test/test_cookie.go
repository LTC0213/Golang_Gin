package main

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) {
	// 获得cookie
	s, err := c.Cookie("username")
	if err != nil {
		s = c.Param("username")
		// 设置cookie
		c.SetCookie("username", s, 60*60, "/", "localhost", false, true)
	}

	c.String(200, "测试cookie\nusername: %s", s)
}

func main() {
	e := gin.Default()
	e.GET("/:username", Handler)
	e.Run()
}
