package main

import . "fmt"
import "math/rand"
import "runtime"
import "runtime/debug"

func main() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	Println(m.Alloc, m.Sys)

	elems := []float64{}

	for i := 0; i < 5e7; i++ {
		elems = append(elems, rand.Float64())
	}

	Println(len(elems))

	m = runtime.MemStats{}
	runtime.ReadMemStats(&m)
	Println(m.Alloc, m.Sys)

	elems = nil
	Println(len(elems))
	runtime.GC()
	debug.FreeOSMemory()

	m = runtime.MemStats{}
	runtime.ReadMemStats(&m)
	Println(m.Alloc, m.Sys)
}
