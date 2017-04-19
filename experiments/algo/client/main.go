package main

import (
	. "../dispatcher"
	. "../types"
	. "fmt"
	"github.com/valyala/gorpc"
	"os"
	"strconv"
	"sync"
)

var servers = []string{
	"localhost:8080",
	"localhost:8081",
	"localhost:8082",
}
var wg sync.WaitGroup
var m1 interface{}
var m2 interface{}

func main() {
	task := os.Getenv("TASK")

	computers, _ := strconv.Atoi(os.Getenv("COMPS"))
	n := int(2)

	switch task {
	case "0":
		m1 = NewPi(n)
	case "1":
		m1, m2 = NewMatrix(n, n, true), NewMatrix(n, n, true)
	case "2":
		m1, m2 = NewMatrix(n, n, true), NewVVector(n, true)
	case "3":
		m1, m2 = NewVVector(n, true), NewHVector(n, true)
	case "4":
		m1 = NewBrute([]string{"95.213.195.96:22"}, []string{"root"}, []string{"Open!3451", "xsgyd12y47"})
	default:
		Println("ALAAAAARM!!!1")
	}

	add, l := 0, n/computers
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

			switch task {
			case "0":
				req := Request0{NewPi(end - start)}
				res, _ := dispclient.Call("PiCalc", &req)
				Println(res.(*Response0).Time)
			case "1":
				req := Request1{m1.(Matrix)[start:end], m2.(Matrix)}
				res, _ := dispclient.Call("MatrixByMatrix", &req)
				Println(res.(*Response1).Time)
			case "2":
				req := Request2{m1.(Matrix)[start:end], m2.(VerticalVector)}
				res, _ := dispclient.Call("MatrixByVVector", &req)
				Println(res.(*Response2).Time)
			case "3":
				req := Request3{m1.(VerticalVector)[start:end], m2.(HorizontalVector)}
				res, _ := dispclient.Call("VVectorByHVector", &req)
				Println(res.(*Response3).Time)
			case "4":
				b := Brute{}
				b.Uds = m1.(Brute).Uds[start:end]
				req := Request4{}
				req.M1 = b
				res, _ := dispclient.Call("Brute", &req)
				Println(res.(*Response4).Time)
				Println(len(res.(*Response4).Answer))
			default:
				Println("NOT FOUND")
			}

		}(i, i*l, i*l+l+add)
	}

	wg.Wait()
}
