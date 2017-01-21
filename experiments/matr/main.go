package main

import (
	// . "fmt"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./main.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "main.html", nil)
	})

	r.GET("/m.js", func(c *gin.Context) {
		c.File("./matreshka.min.js")
	})

	r.Run()
}
