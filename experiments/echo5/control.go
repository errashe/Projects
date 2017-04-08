package main

import . "fmt"
import "net/http"
import "sync"

import "io/ioutil"
import "encoding/json"
import "bytes"

var wg sync.WaitGroup
var servers []string = []string{"http://localhost:8000", "http://localhost:8001"}

func serverAsk(i int, query string) string {
	return Sprintf("%s/%s", servers[i], query)
}

func main() {
	// var n int = 1e6
	var N int = 1

	var m Matrix

	for i := 0; i < 1000; i++ {
		var r Row
		for j := 0; j < 1000; j++ {
			r = append(r, 1000000)
		}
		m = append(m, r)
	}

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			client := &http.Client{}

			req, err := http.NewRequest("GET", serverAsk(i, "settings"), nil)
			req.URL.Query().Add("cores", "2")
			req.URL.Query().Add("threads", "2")

			_, err = client.Do(req)
			if err != nil {
				panic(err)
			}

			body, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}

			req, err = http.NewRequest("POST", serverAsk(i, "run"), bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			req.URL.Query().Add("work", "2")

			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}

			dat, _ := ioutil.ReadAll(resp.Body)
			Println(string(dat))
			// r := gorequest.New()

			// r.Get(Sprintf("%s/settings", servers[i])).
			// 	Param("cores", "2").
			// 	Param("threads", "2").
			// 	End()

			// data, _ := json.Marshal(m)

			// _, body, errs := r.Post(Sprintf("%s/run", servers[i])).
			// 	Type("multipart").
			// 	Param("work", "2").
			// 	Param("par1", Sprintf("%d", n/N)).
			// 	Send(Sprintf("%s", data)).
			// 	EndBytes()

			// if errs != nil {
			// 	Println(errs)
			// } else {
			// 	var res Results
			// 	json.Unmarshal(body, &res)
			// 	Println(res)
			// }
		}(i)
	}

	wg.Wait()
}
