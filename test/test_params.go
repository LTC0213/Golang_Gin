package main

import "github.com/gin-gonic/gin"

func TestQueryString(c *gin.Context) {
	username := c.Query("username")
	site := c.DefaultQuery("site", "www.duoke360.com")

	c.String(200, "username:%s, site:%s", username, site)
}

func main() {

	e := gin.Default()
	// url : http://localhost:8080/testQueryString?username=郭宏志&site=多课网
	e.GET("/testQueryString", TestQueryString)

	e.Run()

}
