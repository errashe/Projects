package main

import . "fmt"
import "net/http"
import "io/ioutil"

import "sync"

// var servers []string = []string{"95.213.194.159", "95.213.195.36", "78.155.206.115", "95.213.251.52"}

var servers []string = []string{"localhost:8000"}

func main() {
	wg := sync.WaitGroup{}

	f := func(url string) {
		defer wg.Done()
		res, _ := http.Get(url)
		resb, _ := ioutil.ReadAll(res.Body)

		Println(string(resb))
	}

	wg.Add(len(servers))
	for i := 0; i < len(servers); i++ {
		go f(Sprintf("http://%s/work?samples=100000000&threads=1&cores=1", servers[i]))
	}

	wg.Wait()
}
