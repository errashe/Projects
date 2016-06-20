package main

import (
	"fmt"
	"net"
	"strings"
)

func fillFunctions() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["QWE"] = func(conn net.Conn, params []string) {
		add := append([]byte{17, 0, 0, 0}, []byte(fmt.Sprintf("Test message with params: %s\n", strings.Join(params, "|")))...)
		conn.Write(add)
	}
	return functions
}
