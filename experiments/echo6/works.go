package main

import (
	"math/rand"
	"time"
)

func work1(res *Results, smpls int, threads int) {
	defer wg.Done()

	t := time.Now()

	samples := smpls / threads

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

	*res = append(*res, Result{ratio * 4, time.Since(t)})
}
