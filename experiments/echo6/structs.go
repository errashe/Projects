package main

import (
	. "fmt"
	"math/rand"
	"time"
)

type Cell float64
type Row []Cell
type Matrix []Row

type Settings struct {
	Cores   int
	Threads int
	Work    int
}

type Data1 struct {
	Num int
}

type Data2 struct {
	A Matrix
	B Matrix
}

type Result struct {
	Data interface{}
	Time time.Duration
}

type Results []Result

func (m *Matrix) fill(n int) {
	for i := 0; i < n; i++ {
		row := Row{}
		for j := 0; j < n; j++ {
			row = append(row, Cell(rand.Float64()))
		}
		*m = append(*m, row)
	}
}

func (m *Matrix) Pretty() {
	for _, row := range *m {
		for _, cell := range row {
			Printf("%.5f ", cell)
		}
		Println()
	}
}
