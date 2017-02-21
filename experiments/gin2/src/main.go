package main

import (
	. "fmt"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	ro *gin.Engine
)

func init() {
	Println("router loaded")

	ro = gin.Default()
}

func main() {

	ro.GET("/", func(c *gin.Context) {
		render.HTML(c.Writer, 200, "main", gin.H{"c": c})
	})

	ro.GET("/set", func(c *gin.Context) {
		s(c).Set("test", "omg")
		s(c).Save()

		c.Redirect(302, "/")
	})

	ro.Run(":8000")

}
