package main

import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"

import . "fmt"
import "time"
import "math/rand"

type Info struct {
	Status        string
	LastWorkValue float64
}

var i Info = Info{"DONE", 0.0}

func main() {
	e := echo.New()

	quit := make(chan bool)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${latency_human}\n",
	}))

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(200, Sprintf("%+v", i))
	})

	e.GET("/run", func(c echo.Context) error {
		if i.Status == "WORK" {
			return c.String(200, "SOMEBODY WORKING NOW")
		}

		go func() {
			i.Status = "WORK"

			t := time.Now()
			cnt := 0.0
			i.LastWorkValue = 0

			for {
				select {
				case <-quit:
					i.LastWorkValue = 0
					i.Status = "DONE"
					return
				default:
					if time.Since(t) < 30*time.Second {
						cnt += 1.0
						i.LastWorkValue += rand.Float64()
					} else {
						i.LastWorkValue /= cnt
						i.Status = "DONE"
						return
					}
				}
			}
		}()

		return c.String(200, "WORKING")
	})

	e.GET("/stop", func(c echo.Context) error {
		quit <- true
		return c.String(200, "STOPPED")
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
