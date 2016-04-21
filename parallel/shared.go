package main

import "sync"
import "math/rand"

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			ch <- rand.Int()
		}
	}()

	go func() {
		defer wg.Done()
		for {
			println(<-ch)
		}
	}()

	wg.Wait()
}
