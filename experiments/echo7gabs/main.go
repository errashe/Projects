package main

import (
	. "fmt"
	"github.com/Jeffail/gabs"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {

		payload := c.FormValue("payload")

		j, _ := gabs.ParseJSON([]byte(payload))
		switch j.Path("work").Data().(float64) {
		case 1:
			Println("FIRST WORK")
		case 2:
			Println("SECOND WORK")
		default:
			Println("WORK 404")
		}

		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
