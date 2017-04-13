package main

import (
	. "fmt"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Fprintf(w, "Hello, from root route!")
	})
}
