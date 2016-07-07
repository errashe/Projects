package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}
	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		log.Fatalf("Invalid port: %s (%s)\n", os.Args[1], err)
	}

	var counter int = 0

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// println("--->", os.Args[1], req.URL.String())
		w.Write([]byte("Hello, world!"))
		counter++
	})

	go func() {
		for range time.Tick(time.Second * 1) {
			println(counter)
		}
	}()

	http.ListenAndServe(":"+os.Args[1], nil)

}
