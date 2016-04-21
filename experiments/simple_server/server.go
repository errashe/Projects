package main

import (
	"log"
	"net"
	"runtime"
	"sync"
	"time"
)

func handle_error(err error, name string) {
	if err != nil {
		log.Println(name, ": ", err.Error())
	}
}

var wg sync.WaitGroup
var err error
var cnt int

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	l, err := net.Listen("tcp", ":1234")
	handle_error(err, "Listen")
	log.Println(l.Addr().String())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range time.Tick(time.Second) {
			log.Println(cnt)
			cnt = 0
		}
	}()

	for {
		conn, err := l.Accept()
		handle_error(err, "Accept")

		cnt++

		conn.Close()
	}

}
