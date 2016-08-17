package main

import (
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com/")
	if err != nil {
		println(err.Error())
	}

	println(res.Status)
}
