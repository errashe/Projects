package types

import (
	"math/rand"
	"time"
)

type Pi struct {
	Samples int
}

func NewPi(n int) Pi {
	return Pi{n}
}

func (p Pi) PiCalc() float64 {
	var inside int = 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < p.Samples; i++ {
		x := r.Float64()
		y := r.Float64()
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	ratio := float64(inside) / float64(p.Samples)

	return ratio * 4
}
