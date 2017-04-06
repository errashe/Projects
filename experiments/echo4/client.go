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
var num int = 1

var results []interface{}

var wg sync.WaitGroup

func work(num, samples int) {
	switch num {
	case 1:
		work1(samples)
	case 2:
		work2(samples)
	}
}

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

func work2(samples int) {
	defer wg.Done()

	results = append(results, float64(samples))
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
		num, _ = strconv.Atoi(r.URL.Query().Get("num"))

		runtime.GOMAXPROCS(cores)

		results = make([]interface{}, 0)

		wg.Add(threads)
		runs := samples / threads
		for i := 0; i < threads; i++ {
			go work(num, runs)
		}
		wg.Wait()

		Fprintf(w, "%+v, %s", results, time.Since(t))
	})

	http.ListenAndServe(":8000", nil)
}
