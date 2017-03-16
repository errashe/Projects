package main

import . "fmt"

type Recepy struct {
	Cnt  float64
	Cost float64
	T1   float64
	T2   float64
	T3   float64
	Ct1  float64
	Ct2  float64
	Ct3  float64
}

func (r *Recepy) cost() float64 {
	t := (*r)
	return t.Cnt * (t.T1*t.Ct1 + t.T2*t.Ct2 + t.T3*t.Ct3)
}

func (r *Recepy) profit() float64 {
	t := (*r)
	return t.Cost * t.Cnt * 1.5
	// ((t.Cnt * .83) +
	// 	(t.Cnt*.20)*2 +
	// 	(t.Cnt*.05)*3 +
	// 	(t.Cnt*.01)*5 +
	// 	(t.Cnt*.005)*10 +
	// 	(t.Cnt*.003)*15 +
	// 	(t.Cnt*.001)*20)
}

func main() {
	grass := Recepy{100, 369, 7, 10, 10, 49.99, 5.87, 8.00}

	Println(grass.cost())
	Println(grass.profit())
	Println(grass.profit() - grass.cost())
}
