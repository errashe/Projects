package main

import (
	. "fmt"
	. "github.com/valyala/gorpc"
	"os"
	"sync"
)

var wgs sync.WaitGroup

func main() {
	wgs.Add(1)

	port := os.Getenv("PORT")
	s := NewTCPServer(Sprintf("127.0.0.1:%s", port), d.NewHandlerFunc())
	if err := s.Start(); err != nil {
		Printf("Cannot start rpc server: [%s]", err)
	}

	wgs.Wait()
}
