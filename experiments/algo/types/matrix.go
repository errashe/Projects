package types

import (
	. "fmt"
	"math/rand"
	"strings"
)

type Matrix []Row

func NewMatrix(i, j int, args ...interface{}) (ret Matrix) {
	random_flag := false
	if len(args) > 0 {
		random_flag = args[0].(bool)
	}

	ret = make(Matrix, i)

	for r := range ret {
		ret[r] = make(Row, j)
		if random_flag {
			for c := range ret[r] {
				ret[r][c] = Cell(rand.Float64())
			}
		}
	}

	return ret
}

func (m Matrix) String() string {
	r := []string{}
	for _, row := range m {
		r = append(r, Sprintf("%s", strings.Join(row.String(), " ")))
	}

	return Sprintf("\r\n%s\r\n", strings.Join(r, "\n"))
}

func (m1 Matrix) MulByMatrix(m2 Matrix) Matrix {

	rows, cols, extra := len(m1), len((m1)[0]), len(m2)
	m3 := make(Matrix, rows)

	for i := 0; i < rows; i++ {
		m3[i] = make(Row, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < extra; k++ {
				m3[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return m3
}

func (m1 Matrix) MulByVVector(m2 VerticalVector) VerticalVector {
	v := NewVVector(len(m2))
	for i := range m1 {
		temp := Cell(0)
		for j := range m1[i] {
			temp += m1[i][j] * m2[j]
		}
		v[i] = Cell(temp)
	}
	return v
}
