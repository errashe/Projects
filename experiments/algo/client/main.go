package main

import (
	. "../dispatcher"
	. "../types"
	. "fmt"
	"github.com/valyala/gorpc"
	"sync"
)

var servers []string = []string{
	"localhost:8080",
	"localhost:8081",
}
var wg sync.WaitGroup

func main() {

	computers := 2
	n := 1000

	m1, m2 := NewMatrix(n, n, true), NewMatrix(n, n, true)
	// req := Request1{NewMatrix(n, n, true), NewMatrix(n, n, true)}
	// res, _ := dispclient.Call("matrix", &req)
	// Println(res)

	req := Request1{NewMatrix(n, n, true), NewMatrix(n, n, true)}

	add, l := 0, len(req.M1)/computers
	for i := 0; i < computers; i++ {
		if i == computers-1 {
			add = n - computers*l
		}
		wg.Add(1)
		go func(i, start, end int) {
			defer wg.Done()

			client := gorpc.NewTCPClient(servers[i])
			client.Start()
			defer client.Stop()
			dispclient := Disp.NewFuncClient(client)

			req := Request1{m1[start:end], m2}

			res, _ := dispclient.Call("matrix", &req)
			Println(res.(*Response1).Time)
		}(i, i*l, i*l+l+add)
	}

	wg.Wait()
}
