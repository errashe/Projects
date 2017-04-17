package main

import (
	. "fmt"
	"github.com/valyala/gorpc"
	"gopkg.in/alecthomas/kingpin.v2"
	"math"
	"sync"
	"time"
)

var wg sync.WaitGroup
var servers []string = []string{
	// "localhost:10001",
	// "localhost:10002",
	"95.213.195.96:12345",
}

var (
	first  = kingpin.Command("first", "First task")
	second = kingpin.Command("second", "Second task")
	third  = kingpin.Command("third", "Third task")
)

func main() {

	servers_active := 1

	switch kingpin.Parse() {
	case "first":
		wg.Add(servers_active)
		for i := 0; i < servers_active; i++ {
			go func(i int) {
				defer wg.Done()

				c := gorpc.NewTCPClient(servers[i])
				c.Start()
				defer c.Stop()
				dc := d.NewFuncClient(c)

				p := Payload1{Data1{1e8 / servers_active}, Settings{2, 2}}

				t := time.Now()
				res, _ := dc.Call("f1", p)
				Println(res, time.Since(t))
			}(i)
		}
		wg.Wait()

	case "second":
		d1, d2 := Matrix{}, Matrix{}
		d1.Fill(1000, 1000)
		d2.Fill(1000, 1000)

		calc := 0
		l := len(d1)
		count := int(math.Ceil(float64(l) / float64(servers_active)))

		for i := 0; i < l; i += count {
			n := i + count
			if n > l {
				n = l
			}

			d1temp := d1[i:n]

			wg.Add(1)
			go func(t1, t2 Matrix, num int) {
				defer wg.Done()

				c := gorpc.NewTCPClient(servers[num])
				c.Start()
				defer c.Stop()
				c.RequestTimeout = time.Minute * 10

				dc := d.NewFuncClient(c)

				p := Payload23{Data23{d1temp, d2}, Settings{2, 2}}

				t := time.Now()
				dc.Call("f2", p)
				Println(time.Since(t))
			}(d1temp, d2, calc)
			calc++
		}

		wg.Wait()
	case "third":
		d1, d2 := Matrix{}, Matrix{}
		d1.Fill(1000, 1)
		d2.Fill(1, 1000)

		calc := 0
		l := len(d1)
		count := int(math.Ceil(float64(l) / float64(servers_active)))

		for i := 0; i < l; i += count {
			n := i + count
			if n > l {
				n = l
			}

			d1temp := d1[i:n]

			wg.Add(1)
			go func(t1, t2 Matrix, num int) {
				defer wg.Done()

				c := gorpc.NewTCPClient(servers[num])
				c.Start()
				defer c.Stop()
				c.RequestTimeout = time.Minute * 10

				dc := d.NewFuncClient(c)

				p := Payload23{Data23{d1temp, d2}, Settings{2, 2}}

				t := time.Now()
				dc.Call("f3", p)
				Println(time.Since(t))
			}(d1temp, d2, calc)
			calc++

			wg.Wait()
		}
	}
}
