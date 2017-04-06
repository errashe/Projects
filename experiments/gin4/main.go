package main

import "github.com/gin-gonic/gin"

type S struct {
	A int `form:"a" json:"a" binding:"required"`
	B int `form:"b" json:"b" binding:"required"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("main.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "main.html", nil)
	})

	r.POST("/some", func(c *gin.Context) {
		f := S{}
		c.BindJSON(&f)
		c.String(200, "%+v", f)
	})

	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
