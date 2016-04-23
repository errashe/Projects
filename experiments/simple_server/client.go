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

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				conn, err := net.Dial("tcp", "localhost:1234")
				handleError(err, "Dial")

				conn.Write(msg)
				conn.Close()
			}
		}()
	}

	wg.Wait()
}
