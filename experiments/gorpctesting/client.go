package main

import "os"
import "log"
import "sync"
import "runtime"
import "github.com/valyala/gorpc"

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 2 {
		log.Fatal("Need specify computer host:port for connecting!")
	}

	for i := 0; i < 100; i++ {
		c := &gorpc.Client{
			// TCP address of the server.
			Addr: os.Args[1],
		}
		c.Start()

		// All client methods issuing RPCs are thread-safe and goroutine-safe,
		// i.e. it is safe to call them from multiple concurrently running goroutines.
		wg.Add(1)
		go func() {
			defer wg.Done()
			for a := 0; a < 10000; a++ {
				resp, err := c.Call("foobar")
				if err != nil {
					log.Fatalf("Error when sending request to server: %s", err)
				}
				if resp.(string) != "foobar" {
					log.Fatalf("Unexpected response from the server: %+v", resp)
				}
			}
		}()
	}
	wg.Wait()
}
