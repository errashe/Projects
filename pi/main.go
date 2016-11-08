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

func PI(samples int) float64 {
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

	return ratio * 4
}

func MultiPI(samples int) float64 {
	cpus := runtime.NumCPU()

	threadSamples := samples / cpus
	results := make(chan float64, cpus)

	for j := 0; j < cpus; j++ {
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
	for i := 0; i < cpus; i++ {
		total += <-results
	}

	return total / float64(cpus)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Using: ./pigo [num of iterations]\n")
		fmt.Printf("Not enough arguments, exiting...\n")
		os.Exit(0)
	}

	count, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("Running Monte-Carlo simulations on %d iterations...\n\n", count)

	t1 := time.Now()
	fmt.Printf("Our value of Pi after \t\t [%d] runs: \t [%.10f] And time: \t [%.5fs] \n",
		count, PI(count), time.Since(t1).Seconds(),
	)

	t2 := time.Now()
	fmt.Printf("Our value of MultiPi after \t [%d] runs: \t [%.10f] And time: \t [%.5fs] \n",
		count, MultiPI(count), time.Since(t2).Seconds(),
	)

	fmt.Printf("Real value of PI: %.10f\n", math.Pi)
}
