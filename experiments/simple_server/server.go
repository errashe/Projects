package main

import (
	"bufio"
	"log"
	"net"
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
	l, err := net.Listen("tcp", ":1234")
	handle_error(err, "Listen")
	log.Println(l.Addr().String())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _ = range time.Tick(time.Second) {
			log.Println(cnt)
			cnt = 0
		}
	}()

	for {
		conn, err := l.Accept()
		handle_error(err, "Accept")

		go func() {
			scanner := bufio.NewScanner(conn)
			for {
				cnt++
				if ok := scanner.Scan(); !ok {
					break
				}
			}

			conn.Close()
		}()
	}

}
