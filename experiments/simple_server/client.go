package main

import (
	"log"
	"net"
	"sync"
)

func handleError(err error, name string) {
	if err != nil {
		log.Fatal(name, ": ", err.Error())
	}
}

var err error
var wg sync.WaitGroup
var msg = []byte("h\n")

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", "localhost:1234")
			handleError(err, "Dial")

			for {
				conn.Write(msg)
			}

			conn.Close()
		}()
	}

	wg.Wait()
}
