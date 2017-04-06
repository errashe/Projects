package main

// import . "fmt"
import "github.com/gin-gonic/gin"

import "time"
import "os/exec"

import "runtime"
import "sync"
import "strconv"

type Settings struct {
	Cores, Threads int
}

type Result struct {
	Res  float64
	Time time.Duration
}

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
	r := gin.Default()

	r.GET("/", root)
	r.GET("run", run)
	r.GET("/settings", settings)
	r.GET("/results", results)

	r.Run(":8000")
}

func root(c *gin.Context) {
	output, _ := exec.Command("uname", "-a").Output()
	c.String(200, "%s", output)
}

func run(c *gin.Context) {
	work, _ := strconv.Atoi(c.Query("work"))
	par1, _ := strconv.Atoi(c.Query("par1"))

	runtime.GOMAXPROCS(s.Cores)

	clearResults()
	wg.Add(s.Threads)
	for i := 0; i < s.Threads; i++ {
		switch work {
		case 1:
			go work1(par1 / s.Threads)
		default:
			c.String(200, "NOTHING")
		}
	}
	wg.Wait()

	c.String(200, "%+v", r)
}

func settings(c *gin.Context) {
	cores, _ := strconv.Atoi(c.Query("cores"))
	threads, _ := strconv.Atoi(c.Query("threads"))

	if cores != 0 {
		s.Cores = cores
	}
	if threads != 0 {
		s.Threads = threads
	}

	c.String(200, "%+v", s)
}

func results(c *gin.Context) {
	c.String(200, "%+v", r)
}
