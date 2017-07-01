package main

import (
	. "fmt"
	"io/ioutil"
	"net/http"
	"time"

	"crypto/md5"
	"encoding/json"

	"github.com/Jeffail/gabs"
)

type Result struct {
	Url   string
	Title string
	Date  time.Time
}

type Results []Result

func Parse() []byte {
	var rs Results
	var res *http.Response
	var err error
	var body []byte

	res, err = http.Get("https://meduza.io/api/v3/search?chrono=news&locale=ru&page=0&per_page=10")
	if err != nil {
		Println(err)
		return nil
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		Println(err)
		return nil
	}

	j, err := gabs.ParseJSON(body)
	if err != nil {
		Println(err)
		return nil
	}

	documents, _ := j.Path("documents").ChildrenMap()

	rs = Results{}

	for _, child := range documents {
		r := Result{}

		r.Url, r.Title, r.Date =
			child.Path("url").Data().(string),
			child.Path("title").Data().(string),
			time.Unix(int64(child.Path("published_at").Data().(float64)), 0)

		rs = append(rs, r)
	}

	rj, _ := json.Marshal(rs)

	return rj
}

func main() {
	md := md5.New()

	var cur_hash []byte = nil
	var cur_res []byte = nil
	var tmp_hash []byte = nil

	for range time.Tick(5 * time.Second) {
		tmp := Parse()
		md.Reset()
		tmp_hash = md.Sum(tmp)

		Printf("%x\n\n", cur_hash)
		Printf("%x\n\n", tmp_hash)

		if Sprintf("%x", cur_hash) != Sprintf("%x", tmp_hash) {
			cur_res = tmp
			cur_hash = tmp_hash

			Printf("%x\n\n%x\n\n", cur_res, cur_hash)
		}
	}
}
