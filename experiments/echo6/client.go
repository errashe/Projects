package main

import (
	"encoding/json"
	. "fmt"
	"net/http"
	"net/url"
	"time"

	"io/ioutil"
)

func main() {
	time.Sleep(1 * time.Second)

	s := Settings{}
	s.Cores = 2
	s.Threads = 2
	s.Work = 1

	d := Data1{}
	d.Num = 1e6

	data1, _ := json.Marshal(s)
	data2, _ := json.Marshal(d)

	vals := url.Values{}

	vals.Add("settings", string(data1))
	vals.Add("data", string(data2))

	resp, err := http.PostForm("http://localhost:1323/work1", vals)
	if err != nil {
		Println(err)
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Println(err)
	}

	Println(string(rbody))
}
