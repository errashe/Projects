package main

import . "fmt"
import . "./types"
import "time"

func main() {
	m := NewMatrix(1000, 1000, true)

	t := time.Now()
	m1, m2 := m[0:500], m[500:1000]
	Println(time.Since(t))

	Println(len(m1), len(m2))

	Printf("%fms\n", 25.0/(1*1024/8)*1000)
}
