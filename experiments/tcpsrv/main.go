package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "192.168.1.31"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var functions map[string]interface{}

func main() {
	functions = fillFunctions()

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024*10)

	l, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	req := string(buf[4 : l-1])
	reqA := strings.Split(req, "|")

	if functions[reqA[0]] != nil {
		functions[reqA[0]].(func(net.Conn, []string))(conn, reqA[1:])
	}
	conn.Close()
}
