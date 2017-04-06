package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
	r.Run(":8000")
}
