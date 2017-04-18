package types

import (
	"math/rand"
	"strings"
)

type Vector interface {
	String()
	NewVector()
}

type VerticalVector []Cell
type HorizontalVector []Cell

func NewVVector(i int, args ...interface{}) (ret VerticalVector) {
	random_flag := false
	if len(args) > 0 {
		random_flag = args[0].(bool)
	}

	ret = make(VerticalVector, i)

	if random_flag {
		for r := range ret {
			ret[r] = Cell(rand.Float64())
		}
	}

	return
}

func NewHVector(i int, args ...interface{}) (ret HorizontalVector) {
	random_flag := false
	if len(args) > 0 {
		random_flag = args[0].(bool)
	}

	ret = make(HorizontalVector, i)

	if random_flag {
		for r := range ret {
			ret[r] = Cell(rand.Float64())
		}
	}

	return
}

func (v1 VerticalVector) MulByHVector(v2 HorizontalVector) (ret Matrix) {
	for i := range v1 {
		ret = append(ret, Row{})
		for j := range v2 {
			ret[i] = append(ret[i], v1[i]*v2[j])
		}
	}

	return
}

func (v VerticalVector) String() string {
	r := []string{}

	for _, cell := range v {
		r = append(r, cell.String())
	}

	return strings.Join(r, " ")
}

func (v HorizontalVector) String() string {
	r := []string{}

	for _, cell := range v {
		r = append(r, cell.String())
	}

	return strings.Join(r, " ")
}
