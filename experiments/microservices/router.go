package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var baseTemplate = "templates/layout.tmpl"

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles(baseTemplate, "templates/post.tmpl")))
		c.HTML(200, "base", nil)
	})

	r.Run(":8000")
}
