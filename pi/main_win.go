package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func MultiPI(samples, cores, threads int) float64 {
	runtime.GOMAXPROCS(cores)

	threadSamples := samples / threads
	results := make(chan float64, threads)

	for j := 0; j < threads; j++ {
		go func() {
			var inside int
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < threadSamples; i++ {
				x, y := r.Float64(), r.Float64()

				if x*x+y*y <= 1 {
					inside++
				}
			}
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	for i := 0; i < threads; i++ {
		total += <-results
	}

	return total / float64(threads)
}

func init() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 2 {
		fmt.Printf("Using: ./pigo [num of iterations]\n")
		fmt.Printf("Not enough arguments, exiting...\n")
		os.Exit(0)
	}
}

func RunMultiPI(count int) {
	cores := []int{1, 2, 3, 4, 5, 6, 7, 8}
	threads := []int{1, 2, 3, 4, 5, 6, 7, 8}

	t := time.Time{}

	for _, core := range cores {
		fmt.Println("Ядер: ", core)
		for _, thread := range threads {
			t = time.Now()
			// fmt.Printf("Run with %d core(s) and %d thread(s)\n", core, thread)
			// fmt.Printf("Our value of MultiPi after \t [%d] runs: \t [%.10f] And time: \t [%.5fs] \n",
			// 	count, MultiPI(count, core, thread), time.Since(t).Seconds(),
			// )

			MultiPI(count, core, thread)
			fmt.Printf("%.10f\n", time.Since(t).Seconds())
		}
	}

}

func main() {
	count, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Running Monte-Carlo simulations on %d iterations...\n\n", count)

	RunMultiPI(count)

	fmt.Printf("Real value of PI: %.10f\n", math.Pi)
}
