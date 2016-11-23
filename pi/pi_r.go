package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

type Work struct {
	Mark  float64
	Value float64
}

var wg sync.WaitGroup
var queue chan (Work)
var results chan (float64)
var current float64 = 4

func main() {
	cpus := runtime.NumCPU()

	runtime.GOMAXPROCS(cpus)
	queue = make(chan Work)
	results = make(chan float64)

	t := time.Now()

	wg.Add(2)
	wg.Add(cpus)

	for i := 0; i < cpus; i++ {
		go func() {
			defer wg.Done()
			for v := range queue {
				results <- v.Mark * 4 / v.Value
			}
		}()
	}

	go func() {
		defer wg.Done()
		value := 3.0
		mark := -1.0
		for {
			wrk := Work{}
			wrk.Mark = mark
			wrk.Value = value

			queue <- wrk

			value += 2
			mark = -mark
		}
	}()

	go func() {
		defer wg.Done()

		go func() {
			for {
				accuracy := math.Abs(math.Pi - current)
				fmt.Printf("%.30f accuracy\n", accuracy)
				if accuracy <= 0.00000001 {
					fmt.Println("Результат достигнут, точность 8 знаков получена")

					fmt.Println(time.Since(t))
					os.Exit(0)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()

		for res := range results {
			current += res
		}
	}()

	wg.Wait()
}
