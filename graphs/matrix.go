package main

import (
	"fmt"
	"math"
)

type MiniPair struct {
	x, y int
}

type Cell struct {
	pos          MiniPair
	availability bool
	dist         int
	koef         int
}

type Row []Cell
type Matrix [][]Cell

func (m *Matrix) FillEmptyMatrix(cells int) {
	for i := 0; i < cells; i++ {
		(*m) = append(*m, Row{})
		for j := 0; j < cells; j++ {
			(*m)[i] = append((*m)[i], Cell{MiniPair{0, 0}, true, 0, 0})
		}
	}
}

func (m *Matrix) FillByPairs(pairs Pairs) {
	for _, pair := range pairs {
		(*m)[pair.first.Id][pair.second.Id].dist = int(math.Ceil(pair.dist))
		(*m)[pair.first.Id][pair.second.Id].pos = MiniPair{pair.first.Id, pair.second.Id}
		if pair.first.Id == pair.second.Id {
			(*m)[pair.first.Id][pair.second.Id].availability = false
		}
	}
}

func (m *Matrix) Print() {
	for _, row := range *m {
		for _, cell := range row {
			if cell.availability {
				fmt.Printf("%d(%d)\t", cell.dist, cell.koef)
			} else {
				fmt.Printf("M(0)\t")
			}
		}
		fmt.Printf("\n")
	}
}

func (m *Matrix) RowMin(i int, exclude ...MiniPair) int {
	var min int = 10e10

	for j := 0; j < len((*m)); j++ {
		if len(exclude) > 0 {
			if exclude[0].x == i && exclude[0].y == j {
				continue
			}
		}
		if !(*m)[i][j].availability {
			continue
		}
		if (*m)[i][j].dist < min {
			min = (*m)[i][j].dist
		}
	}

	return min
}

func (m *Matrix) ReduxRow(i int) {
	v := m.RowMin(i)
	for j := 0; j < len((*m)[0]); j++ {
		if (*m)[i][j].availability {
			(*m)[i][j].dist -= v
		}
	}
}

func (m *Matrix) ReduxRows() {
	for i := range *m {
		m.ReduxRow(i)
	}
}

func (m *Matrix) ColMin(i int, exclude ...MiniPair) int {
	var min int = 10e10

	for j := 0; j < len((*m)); j++ {
		if len(exclude) > 0 {
			if exclude[0].x == j && exclude[0].y == i {
				continue
			}
		}
		if !(*m)[j][i].availability {
			continue
		}
		if (*m)[j][i].dist < min {
			min = (*m)[j][i].dist
		}
	}

	return min
}

func (m *Matrix) ReduxCol(i int) {
	v := m.ColMin(i)

	for j := 0; j < len((*m)); j++ {
		(*m)[j][i].dist -= v
	}
}

func (m *Matrix) ReduxCols() {
	for i := range *m {
		m.ReduxCol(i)
	}
}

func (m *Matrix) CalcKoef() {
	for i := range *m {
		for j := range (*m)[i] {
			if (*m)[i][j].availability && (*m)[i][j].dist == 0 {
				(*m)[i][j].koef = m.RowMin(i, MiniPair{i, j}) + m.ColMin(j, MiniPair{i, j})
			}
		}
	}
}

func (m *Matrix) FindMaxKoef() []MiniPair {
	max := 0
	for i := range *m {
		for j := range (*m)[i] {
			if (*m)[i][j].availability && max < (*m)[i][j].koef {
				max = (*m)[i][j].koef
			}
		}
	}

	var res []MiniPair

	for i := range *m {
		for j := range (*m)[i] {
			if (*m)[i][j].koef == max {
				el := (*m)[i][j].pos
				res = append(res, MiniPair{el.x, el.y})
			}
		}
	}

	return res
}

func (m *Matrix) FindRow() int {
	for i := range *m {
		var found bool = false
		for j := range (*m)[i] {
			if !(*m)[i][j].availability {
				found = true
			}
		}
		if !found {
			return i
		}
	}
	return -1
}

func (m *Matrix) FindCol() int {
	for i := range *m {
		var found bool = false
		for j := range (*m)[i] {
			if !(*m)[j][i].availability {
				found = true
			}
		}
		if !found {
			return i
		}
	}
	return -1
}

func (m *Matrix) FindAndClean() {
	row := m.FindRow()
	col := m.FindCol()
	if row != -1 && col != -1 {
		(*m)[row][col].availability = false
	}
}

func (m *Matrix) DeleteRowCol(row, col int) {
	m.DeleteRow(row)
	m.DeleteCol(col)
	m.FindAndClean()
}

func (m *Matrix) DeleteRow(row int) {
	rm := Matrix{}
	rm.FillEmptyMatrix(len((*m)) - 1)

	a := 0
	for i := range *m {
		if i != row {
			rm[a] = (*m)[i]
			a++
		}
	}

	*m = rm
}

func (m *Matrix) DeleteCol(col int) {
	rm := Matrix{}
	rm.FillEmptyMatrix(len((*m)))

	for i := range *m {
		a := 0
		for j := range (*m)[i] {
			if j != col {
				rm[i][a] = (*m)[i][j]
				a++
			}
		}
	}

	*m = rm
}
