package main

import . "fmt"
import "github.com/labstack/echo"

import "os"
import "os/exec"

import "runtime"
import "sync"

var (
	s Settings
	r []Result

	wg sync.WaitGroup
)

func init() {
	s = Settings{1, 1}
	// r = Result{0, time.Since(time.Now())}
}

func clearResults() {
	r = nil
}

func addResult(res Result) {
	r = append(r, res)
}

func main() {
	r := echo.New()

	r.GET("/", root)
	r.POST("run", run)
	r.GET("/settings", settings)
	r.GET("/results", results)

	port := os.Getenv("PORT")
	r.Logger.Fatal(r.Start(Sprintf(":%s", port)))
}

func root(c echo.Context) error {
	output, _ := exec.Command("uname", "-a").Output()
	return c.String(200, Sprintf("%s", output))
}

func run(c echo.Context) error {
	var rp ReqParams
	c.Bind(&rp)

	runtime.GOMAXPROCS(s.Cores)

	clearResults()
	wg.Add(s.Threads)
	for i := 0; i < s.Threads; i++ {
		switch rp.Work {
		case 1:
			go work1(rp)
		case 2:
			go work2(rp)
		default:
			wg.Add(-s.Threads)
			return c.String(200, "NOTHING")
		}
	}
	wg.Wait()

	return c.JSON(200, r)
}

func settings(c echo.Context) error {
	c.Bind(&s)

	return c.String(200, Sprintf("%+v", s))
}

func results(c echo.Context) error {
	return c.JSON(200, r)
}
