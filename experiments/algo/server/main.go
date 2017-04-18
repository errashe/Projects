package main

import (
	. "../dispatcher"
	. "fmt"
	"github.com/valyala/gorpc"
	"os"
)

func main() {
	server := gorpc.NewTCPServer(Sprintf(":%s", os.Getenv("PORT")), Disp.NewHandlerFunc())

	if err := server.Serve(); err != nil {
		Println(err)
	}
}
