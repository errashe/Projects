package main

import . "github.com/valyala/gorpc"
import "time"
import "math"

var servers []string = []string{
	"127.0.0.1:10001",
	"127.0.0.1:10002",
	"127.0.0.1:10003",
	"127.0.0.1:10004",
}

func main() {
	d1, d2 := Matrix{}, Matrix{}
	d1.Fill(1000)
	d2.Fill(1000)

	computers := 4
	l := len(d1) - 1

	calc := 0
	count := int(math.Ceil(float64(l) / float64(computers)))
	for i := 0; i < l; i += count {
		n := i + count
		if n > l {
			n = l
		}

		d1temp := d1[i:n]

		wg.Add(1)
		go func(t1, t2 Matrix, num int) {
			defer wg.Done()

			c := NewTCPClient(servers[num])
			c.Start()
			defer c.Stop()
			c.RequestTimeout = time.Minute * 10

			dc := d.NewFuncClient(c)

			pl := Payload1{}
			pl.S.Cores = 2
			pl.S.Threads = 2
			pl.D.M1 = d1temp
			pl.D.M2 = d1temp

			dc.Call("f2", &pl)
		}(d1temp, d2, calc)
		calc++
	}

	wg.Wait()

}
