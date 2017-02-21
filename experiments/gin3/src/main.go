package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/unrolled/render"
	"gopkg.in/gin-gonic/gin.v1"

	// . "fmt"
	"html/template"
)

var (
	r       *gin.Engine
	options render.Options
	funcs   template.FuncMap
	html    *render.Render
	h       func(c *gin.Context, name string, params gin.H)
	store   sessions.CookieStore
	s       func(c *gin.Context) sessions.Session
)

func getSession(c *gin.Context) int {
	tmp := s(c).Get("test")

	if tmp != nil {
		return tmp.(int)
	} else {
		return 0
	}
}

func init() {
	r = gin.Default()

	html = render.New(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl"},
		Funcs: []template.FuncMap{template.FuncMap{
			"getSession": getSession,
		}},
		Delims:        render.Delims{"${", "}"},
		IsDevelopment: true,
	})

	h = func(c *gin.Context, name string, params gin.H) {
		params["c"] = c
		html.HTML(c.Writer, 200, name, params)
	}

	store = sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("gin", store))

	s = func(c *gin.Context) sessions.Session { return sessions.Default(c) }
}

func main() {
	r.GET("/", rootPath)

	pages := r.Group("/pages")
	pages.GET("/:id/get", func(c *gin.Context) { c.String(200, "current page") })

	adminpages := r.Group("/admin")
	adminpages.GET("/", func(c *gin.Context) { c.String(200, "admin / page") })
	adminpages.GET("/create", func(c *gin.Context) { c.String(200, "admin/create") })
	adminpages.GET("/:id/edit", func(c *gin.Context) { c.String(200, "admin/edit page") })
	adminpages.POST("/:id/save", func(c *gin.Context) { c.String(200, "admin/save page") })
	adminpages.GET("/:id/delete", func(c *gin.Context) { c.String(200, "admin/delete page") })

	r.Run(":8000")
}

func rootPath(c *gin.Context) {
	h(c, "main", gin.H{"qwe": "test"})
}
