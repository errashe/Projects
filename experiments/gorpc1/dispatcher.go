package main

import (
	. "fmt"
	. "github.com/valyala/gorpc"
	"math"
	"runtime"
	"strings"
	"sync"
	"time"
)

var d Dispatcher
var wg sync.WaitGroup

func init() {
	d = *NewDispatcher()

	d.AddFunc("f2", func(p *Payload1) string {

		results := make([]string, p.S.Threads)

		runtime.GOMAXPROCS(p.S.Cores)

		wg.Add(p.S.Threads)
		t := time.Now()

		l := len(p.D.M1) - 1
		calc := 0
		count := int(math.Ceil(float64(l) / float64(p.S.Threads)))
		for i := 0; i < l; i += count {
			n := i + count
			if n > l {
				n = l
			}
			// Println(q[i:n])
			M1temp := p.D.M1[i:n]
			go func(t1, t2 Matrix, calc int) {
				defer wg.Done()

				m3 := t1.MulBy(&t2)
				results[calc] = Sprintf("%v\n", m3.PP())
			}(M1temp, p.D.M2, calc)
			calc++
		}

		wg.Wait()
		Println(time.Since(t))

		return strings.Join(results, "\n")
	})
}
