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

func (p *Points) Find(i int) (float64, float64) {
	for _, e := range *p {
		if e.Id == i {
			return e.X, e.Y
		}
	}
	return 0, 0
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

func (p *Pair) Reverse() Pair {
	temp := Pair{}
	temp.first = p.second
	temp.second = p.first
	temp.dist = p.dist
	return temp
}

func Grad2Rad(grad float64) float64 {
	return grad * math.Pi / 180
}

func (p *Pair) CalcDist() {
	temp := math.Sin(Grad2Rad(p.first.X))*math.Sin(Grad2Rad(p.second.X)) + math.Cos(Grad2Rad(p.first.X))*math.Cos(Grad2Rad(p.second.X))*math.Cos(Grad2Rad(p.first.Y)-Grad2Rad(p.second.Y))
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
