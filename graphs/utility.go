package main

import (
	"math"
)

type Point struct {
	X, Y       float64
	Name, Hash string
	Id         int
}

type Points []Point

func (p *Points) Find(i int) Point {
	for _, e := range *p {
		if e.Id == i {
			return e
		}
	}
	return Point{}
}

func (p *Points) Merge(ip Points) {
	for _, i := range ip {
		if !(*p).Isset(i) {
			(*p) = append((*p), i)
		}
	}
}

func (p *Points) Isset(ip Point) bool {
	for _, i := range *p {
		if i.Hash == ip.Hash {
			return true
		}
	}
	return false
}

func (p *Points) Numerate() {
	for i := range *p {
		(*p)[i].Id = i
	}
}

type Pair struct {
	first, second Point
	dist          float64
}

func Grad2Rad(grad float64) float64 {
	return grad * math.Pi / 180
}

func (p *Pair) CalcDist() {
	p1 := math.Sin(Grad2Rad(p.first.X)) * math.Sin(Grad2Rad(p.second.X))
	p2 := math.Cos(Grad2Rad(p.first.X)) * math.Cos(Grad2Rad(p.second.X))
	p3 := math.Cos(Grad2Rad(p.first.Y) - Grad2Rad(p.second.Y))
	temp := p1 + p2*p3
	temp = math.Acos(temp)
	p.dist = temp * 6371000
}

type Pairs []Pair

func (p *Pairs) Isset(ip Pair) bool {
	for _, i := range *p {
		if i == ip {
			return true
		}
	}
	return false
}

func (p *Pairs) FillByPoints(po *Points) {
	for i := range *po {
		for j := range *po {
			pair := Pair{}
			pair.first = (*po)[i]
			pair.second = (*po)[j]
			pair.CalcDist()
			*p = append(*p, pair)
		}
	}
}
