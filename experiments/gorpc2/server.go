package main

import (
	. "fmt"
	"github.com/valyala/gorpc"
	"os"
)

func main() {
	server := gorpc.NewTCPServer(Sprintf(":%s", os.Getenv("PORT")), d.NewHandlerFunc())
	if err := server.Serve(); err != nil {
		Println(err)
	}
}
