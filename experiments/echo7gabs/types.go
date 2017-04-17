package main

import (
	. "fmt"
	"math/rand"
	"strings"
)

type Row []float64
type Matrix []Row

func (m *Matrix) Fill(n int) {
	for i := 0; i < n; i++ {
		r := Row{}
		for j := 0; j < n; j++ {
			r = append(r, rand.Float64())
		}
		*m = append(*m, r)
	}
}

func (m *Matrix) PP() string {
	var gtemp []string
	for _, row := range *m {
		var temp []string
		for _, cell := range row {
			temp = append(temp, Sprintf("%f", cell))
		}
		gtemp = append(gtemp, strings.Join(temp, " "))
	}

	return Sprintf("%s", strings.Join(gtemp, "\n"))
}

func (m1 *Matrix) MulBy(m2 *Matrix) *Matrix {
	m11, m22 := *m1, *m2

	rows, cols, extra := len(*m1), len((*m1)[0]), len(*m2)
	m3 := make(Matrix, rows)

	for i := 0; i < rows; i++ {
		m3[i] = make(Row, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < extra; k++ {
				m3[i][j] += m11[i][k] * m22[k][j]
			}
		}
	}

	return &m3
}

type Settings struct {
	Cores, Threads int
}

type Module1 struct {
	A, B Matrix
}
