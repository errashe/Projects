package main

import (
	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/unrolled/render"

	. "fmt"
	"html/template"
	"math/rand"
	"time"
)

var (
	r       *echo.Echo
	options render.Options
	funcs   template.FuncMap
	html    *render.Render
	h       func(c echo.Context, name string, params map[string]interface{}) error
	store   session.CookieStore
	s       func(c echo.Context) session.Session
)

func getSession(c echo.Context) (res int) {
	tmp := s(c).Get("test")

	if tmp != nil {
		res = tmp.(int)
	}
	return
}

func init() {
	r = echo.New()

	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} |${status}| ${latency_human} | ${method} ${path}",
	}))

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

	h = func(c echo.Context, name string, params map[string]interface{}) error {
		if params == nil {
			params = make(map[string]interface{})
		}

		params["c"] = c
		return html.HTML(c.Response().Writer, 200, name, params)
	}

	store = session.NewCookieStore([]byte("secret"))
	r.Use(session.Sessions("echo", store))

	s = func(c echo.Context) session.Session { return session.Default(c) }

	rand.Seed(time.Now().UnixNano())
}

func main() {
	r.GET("/", rootPath)

	r.GET("/add", func(c echo.Context) error {
		s(c).Set("test", rand.Int())
		s(c).Save()

		return c.Redirect(302, "/")
	})

	r.GET("/del", func(c echo.Context) error {
		s(c).Delete("test")
		s(c).Save()

		return c.Redirect(302, "/")
	})

	pages := r.Group("/pages")
	pages.GET("/", func(c echo.Context) error { return c.String(200, "pages / page") })
	pages.GET("/:id", func(c echo.Context) error { return c.String(200, Sprintf("pages /%s page", c.Param("id"))) })

	adminpages := pages.Group("/admin")
	adminpages.GET("/", func(c echo.Context) error { return c.String(200, "admin / page") })
	adminpages.GET("/create", func(c echo.Context) error { return c.String(200, "admin/create") })
	adminpages.GET("/:id/edit", func(c echo.Context) error { return c.String(200, "admin/edit page") })
	adminpages.POST("/:id/save", func(c echo.Context) error { return c.String(200, "admin/save page") })
	adminpages.GET("/:id/delete", func(c echo.Context) error { return c.String(200, "admin/delete page") })

	r.Start(":8000")
}

func rootPath(c echo.Context) error {
	return h(c, "main", nil)
}
