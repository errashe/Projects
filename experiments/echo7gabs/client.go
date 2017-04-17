package main

import (
	"encoding/json"
	. "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	s := Settings{2, 2}

	m := Module1{}

	a, b := Matrix{}, Matrix{}
	a.Fill(300)
	b.Fill(300)

	t := time.Now()
	a.MulBy(&b)
	Println(time.Since(t))
	// Println(ss.PP())

	m.A = a
	m.B = b

	data, _ := json.Marshal(m)
	data2, _ := json.Marshal(s)

	resp, _ := http.PostForm("http://localhost:1323/work?uid=1", url.Values{
		"payload":  {string(data)},
		"settings": {string(data2)},
	})

	rdata, _ := ioutil.ReadAll(resp.Body)
	var temp interface{}
	json.Unmarshal(rdata, &temp)

	// Println()
	// for _, row := range temp.([]interface{}) {
	// 	Printf("%s", row)
	// }
}
