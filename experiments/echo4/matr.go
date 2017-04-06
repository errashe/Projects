package main

import . "fmt"
import "math/rand"
import "time"

type Cell float64
type Row []Cell
type Matrix []Row

func (m *Matrix) String() {
	for _, row := range *m {
		Print(row, "\n")
	}
}

func makeMatr(num int) Matrix {
	A := make(Matrix, num)

	for i := 0; i < num; i++ {
		A[i] = make(Row, num)
		for j := 0; j < num; j++ {
			A[i][j] = Cell(rand.Float64())
		}
	}

	return A
}

func muler(A, B Matrix) Matrix {
	C := make(Matrix, len(A))

	for i := 0; i < len(A); i++ {
		C[i] = make(Row, len(A))
		for j := 0; j < len(A); j++ {
			C[i][j] = 0
			for k := 0; k < len(A); k++ {
				C[i][j] = C[i][j] + (A[i][k] * B[k][j])
			}
		}
	}

	return C
}

func main() {
	t := time.Now()

	A := makeMatr(16)
	B := makeMatr(16)

	C := muler(A, B)

	C.String()
	Println(time.Since(t))
}
