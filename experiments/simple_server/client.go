package main

import (
	"log"
	"net"
	"runtime"
	"sync"
)

func handleError(err error, name string) {
	if err != nil {
		log.Fatal(name, ": ", err.Error())
	}
}

var err error
var wg sync.WaitGroup
var msg = []byte("Hello")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for th := 0; th < 100; th++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				conn, err := net.Dial("tcp", "95.213.195.152:1234")
				handleError(err, "Dial")

				conn.Write(msg)
				conn.Close()
			}
		}()
	}

	wg.Wait()
}
