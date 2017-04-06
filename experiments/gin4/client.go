package main

import . "fmt"
import "net/http"
import "io/ioutil"
import "bytes"

import "github.com/Jeffail/gabs"

func main() {
	json := gabs.New()

	json.SetP(10, "a")
	json.SetP(20, "b")

	res, _ := http.Post("http://localhost:8080/some", "application/json", bytes.NewBuffer(json.Bytes()))

	data, _ := ioutil.ReadAll(res.Body)
	Printf("%s", data)
}
