package main

import (
	. "fmt"
	"net/http"
	"os"
	"sync"
)

var port string
var wg sync.WaitGroup

func init() {
	if port = Sprintf(":%s", os.Getenv("PORT")); port == ":" {
		port = ":8000"
	}
}

func main() {

	wg.Add(1)
	go func() {
		defer wg.Done()
		Println(http.ListenAndServe(port, nil))
	}()
	Printf("Server started at %s port\n", port)
	wg.Wait()
}
