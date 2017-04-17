package main

import . "fmt"
import "math"

func main() {
	q := make([]int, 133)

	l := len(q)
	count := int(math.Ceil(float64(l) / 2))
	for i := 0; i < l; i += count {
		n := i + count
		if n > l {
			n = l
		}
		Println(q[i:n])
	}
}
