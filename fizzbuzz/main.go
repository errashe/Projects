package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		var str string = fmt.Sprintf("%d", i)

		if i%3 == 0 {
			str = "fizz"
		}
		if i%5 == 0 {
			str = "buzz"
		}
		if i%15 == 0 {
			str = "fizzbuzz"
		}

		fmt.Println(str)
	}
}
