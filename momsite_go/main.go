package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLFiles("views/main.html")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "main.html", nil)
	// })

	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "main.html", nil)
	})

	r.Run(":3000")
}
