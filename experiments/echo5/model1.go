package main

import "time"
import "math/rand"

// import . "fmt"

func work1(rp ReqParams) {
	defer wg.Done()

	samples := rp.Par1 / s.Threads

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

	addResult(Result{ratio * 4, time.Since(t)})
}
