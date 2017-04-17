package main

import (
	. "fmt"
	"github.com/valyala/gorpc"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Settings struct {
	Cores, Threads int
}

type Row []float64
type Matrix []Row

type Data23 struct {
	M1, M2 Matrix
}

type Payload23 struct {
	D Data23
	S Settings
}

type Data1 struct {
	Num int
}

type Payload1 struct {
	D Data1
	S Settings
}

var wgs sync.WaitGroup
var d *gorpc.Dispatcher

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	d = gorpc.NewDispatcher()

	d.AddFunc("f3", func(p *Payload23) {
		p.D.M1.MulByVector(&p.D.M2)
	})

	d.AddFunc("f2", func(p *Payload23) {
		t := time.Now()
		p.D.M1.MulByMatrix(&p.D.M2)
		Println(time.Since(t))
	})

	d.AddFunc("f1", func(p *Payload1) []float64 {
		var inside int = 0
		var results []float64

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < p.D.Num; i++ {
			x := r.Float64()
			y := r.Float64()
			if (x*x + y*y) < 1 {
				inside++
			}
		}

		ratio := float64(inside) / float64(p.D.Num)
		results = append(results, ratio*4)

		return results
	})
}

func (m *Matrix) Fill(rc int, cc int) {
	for i := 0; i < rc; i++ {
		r := Row{}
		for j := 0; j < cc; j++ {
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

func (m1 *Matrix) MulByMatrix(m2 *Matrix) *Matrix {
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

func (m1 *Matrix) MulByVector(m2 *Matrix) {
	m11, m22 := *m1, *m2

	m3 := Matrix{}

	for i := 0; i < len(m11); i++ {
		r := Row{}
		for j := 0; j < len(m11); j++ {
			r = append(r, m11[i][0]*m22[0][j])
		}
		m3 = append(m3, r)
	}

	Println(m3.PP())
}
