package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/unrolled/render"
	"gopkg.in/olahol/melody.v1"
)

var e *echo.Echo
var r *render.Render
var m *melody.Melody

func init() {
	e = echo.New()

	r = render.New(render.Options{
		IsDevelopment: true,
	})

	m = melody.New()

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
}

func main() {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${method} ${uri} ${status}\n"}))

	e.GET("/", func(c echo.Context) error {
		return r.HTML(c.Response().Writer, 200, "index", nil)
	})

	e.GET("/ws", func(c echo.Context) error {
		return m.HandleRequest(c.Response().Writer, c.Request())
	})

	e.Logger.Fatal(e.StartTLS(":8080", "server.crt", "server.key"))
}
