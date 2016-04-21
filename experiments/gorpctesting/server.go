package main

import "log"
import "sync"
import "time"
import "runtime"
import "github.com/valyala/gorpc"

var cnt int
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range time.Tick(time.Second) {
			log.Println(cnt, x)
			cnt = 0
		}
	}()

	s := &gorpc.Server{
		// Accept clients on this TCP address.
		Addr: ":1234",

		// Echo handler - just return back the message we received from the client
		Handler: func(clientAddr string, request interface{}) interface{} {
			// log.Printf("Obtained request %+v from the client %s\n", request, clientAddr)
			cnt++
			return request
		},
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc server: %s", err)
	}

	wg.Wait()

}
