package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

var results = make(chan int)
var aveResults = make(chan float64)

func calc() {
	defer wg.Done()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		x, y := r.Float64(), r.Float64()
		if x*x+y*y <= 1 {
			results <- 1
		} else {
			results <- 0
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(runtime.NumCPU() + 2)

	for i := 0; i < runtime.NumCPU(); i++ {
		go calc()
	}

	go func() {
		defer wg.Done()

		var validCounter float64 = 0
		var allCounter float64 = 0

		for val := range results {
			if val == 1 {
				validCounter++
			}
			allCounter++
			aveResults <- 4 * validCounter / allCounter
		}
	}()

	go func() {
		defer wg.Done()

		var currentPi float64

		for {
			go func() {
				for res := range aveResults {
					currentPi = res
				}
			}()

			<-time.After(1 * time.Second)
			fmt.Printf("Current result of pi: %.20f\n", currentPi)

		}
	}()

	wg.Wait()
}
