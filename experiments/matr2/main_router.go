package main

import (
	"github.com/gin-gonic/gin"
)

func main_router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("./main.html")

	r.GET("/*all", func(c *gin.Context) {
		c.HTML(200, "main.html", nil)
	})

	return r
}
