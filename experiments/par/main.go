package main

import . "fmt"
import "net/http"
import "os/exec"
import "strconv"
import "sync"
import "runtime"

type Settings struct {
	Cores   int
	Threads int
}

var s Settings
var wgMain sync.WaitGroup
var wg sync.WaitGroup
var results []interface{}

func init() {
	s = Settings{1, 1}
}

func clearResults() {
	results = make([]interface{}, 0)
}

func addResult(item interface{}) {
	results = append(results, item)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("uname", "-a")
		out, _ := cmd.Output()

		Fprintf(w, "%s", out)
	})

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		work, _ := strconv.Atoi(r.URL.Query().Get("work"))
		par1, _ := strconv.Atoi(r.URL.Query().Get("par1"))

		runtime.GOMAXPROCS(s.Cores)

		clearResults()
		wg.Add(s.Threads)
		for i := 0; i < s.Threads; i++ {
			switch work {
			case 1:
				go work1(par1 / s.Threads)
			default:
				Println("NOTHING")
			}
		}
		wg.Wait()

		Fprintf(w, "%+v", results)
	})

	http.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		Fprintf(w, "%+v", results)
	})

	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		cores, _ := strconv.Atoi(r.URL.Query().Get("cores"))
		threads, _ := strconv.Atoi(r.URL.Query().Get("threads"))
		if cores != 0 {
			s.Cores = cores
		}
		if threads != 0 {
			s.Threads = threads
		}
		Fprintf(w, "%+v", s)
	})

	wgMain.Add(1)
	go func() {
		defer wgMain.Done()
		http.ListenAndServe(":8000", nil)
	}()
	Println("Server started on :8000 port")

	wgMain.Wait()
}
