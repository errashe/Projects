package main

import (
	"encoding/json"
	. "fmt"
	"runtime"
	"sync"

	"github.com/labstack/echo"
)

var wg sync.WaitGroup

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Some stats here")
	})

	e.POST("/work1", func(c echo.Context) error {
		s := Settings{}
		d := Data1{}
		var res Results

		json.Unmarshal([]byte(c.FormValue("settings")), &s)
		json.Unmarshal([]byte(c.FormValue("data")), &d)

		runtime.GOMAXPROCS(s.Cores)
		wg.Add(s.Threads)

		for i := 0; i < s.Threads; i++ {
			go work1(&res, d.Num, s.Threads)
		}

		wg.Wait()
		return c.String(200, Sprintf("%+v\n\n%+v\n\n%+v", s, d, res))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
