package main

import (
	. "fmt"
	"hash/adler32"
	"time"
)

func main() {
	t := time.Now()

	for run := 0; run < 3; run++ {
		for i := 0; i < 5e6; i++ {
			adler32.Checksum([]byte("Какая-то не очень длинная строка, для которой надо посчитать контрольную сумму"))
		}
	}

	Println(time.Since(t) / 3)
}
