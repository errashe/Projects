package main

import "fmt"

func main() {
	num := 100
	ch := make(chan int, num)

	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)

	for q := range ch {
		fmt.Println(q)
	}

}
