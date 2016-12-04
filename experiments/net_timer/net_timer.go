package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{}
	client.Timeout = time.Second * 2
	status := make(chan string)

	start := time.Now()

	go func() {
		for {
			_, err := client.Get("http://ya.ru/")
			if err == nil {
				status <- "INTERNET"
				break
			}
			status <- "NOPE"
		}
	}()

	for s := range status {
		fmt.Printf("[%s]: %s\n", time.Since(start), s)
		if s == "INTERNET" {
			os.Exit(0)
		}
	}
}
