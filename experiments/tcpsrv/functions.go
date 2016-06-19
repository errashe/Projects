package main

import (
	"fmt"
	"net"
	"strings"
)

func fillFunctions() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["WAR"] = func(conn net.Conn, params []string) {
		conn.Write([]byte(fmt.Sprintf("Test message with params: %s\n", strings.Join(params, "|"))))
	}
	return functions
}
