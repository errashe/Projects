package main

import "time"
import "math/rand"
import . "fmt"

type Res struct {
	Result interface{}
	time   string
}

func work1(samples int) {
	defer wg.Done()
	t := time.Now()
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

	addResult(Res{ratio * 4, Sprintf("%s", time.Since(t))})
}
