package main

import . "fmt"
import "net/http"
import "runtime"
import "strconv"

import "time"
import "math/rand"
import "sync"

import "os/exec"

var cores int = 1
var threads int = 1
var samples int = 1e6

var results []interface{}

var wg sync.WaitGroup

func work1(samples int) {
	defer wg.Done()
	var inside int = 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < samples; i++ {
		x := r.Float64()
		y := r.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	ratio := float64(inside) / float64(samples)

	results = append(results, ratio*4)
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out, _ := exec.Command("uname", "-a").Output()
		Fprint(w, string(out))
	})

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		cores, _ = strconv.Atoi(r.URL.Query().Get("cores"))
		threads, _ = strconv.Atoi(r.URL.Query().Get("threads"))
		samples, _ = strconv.Atoi(r.URL.Query().Get("samples"))

		runtime.GOMAXPROCS(cores)

		results = make([]interface{}, 0)

		wg.Add(threads)
		for i := 0; i < threads; i++ {
			go work1(samples)
		}
		wg.Wait()

		res := 0.0
		for _, val := range results {
			res += val.(float64)
		}

		Fprintf(w, "%+v, %.5f, %s", results, res/float64(threads), time.Since(t))
	})

	http.ListenAndServe(":8000", nil)
}
