package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/unrolled/render"
	"gopkg.in/gin-gonic/gin.v1"
)

func s(c *gin.Context) sessions.Session { return sessions.Default(c) }

func main() {
	defer session.Close()
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	h := render.New(render.Options{
		Layout:        "layout",
		IsDevelopment: true,
		Funcs:         funcs,
	})

	r.GET("/", func(c *gin.Context) {
		h.HTML(c.Writer, 200, "main", gin.H{"c": c})
	})

	r.GET("/login", func(c *gin.Context) {
		s(c).Set("user", "e4stw00d")
		s(c).Save()

		c.Redirect(302, "/")
	})

	r.GET("/logout", func(c *gin.Context) {
		s(c).Delete("user")
		s(c).Save()

		c.Redirect(302, "/")
	})

	r.Run(":8000")
}
