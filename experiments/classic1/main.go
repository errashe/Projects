package main

import (
	. "fmt"
	// "net/http"
	"os"
)

var port string

func main() {
	if port = Sprintf(":%s", os.Getenv("PORT")); port == ":" {
		port = ":8000"
	}
	Println(port)
}
