package main

import (
	"encoding/json"
	. "fmt"
	"math"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var results []string
var wg sync.WaitGroup

func main() {
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))

		payload := r.FormValue("payload")
		var settings Settings
		json.Unmarshal([]byte(r.FormValue("settings")), &settings)
		results = make([]string, settings.Threads)

		runtime.GOMAXPROCS(settings.Cores)

		switch uid {
		case 1:
			var m Module1
			json.Unmarshal([]byte(payload), &m)

			wg.Add(settings.Threads)

			l := len(m.A)
			calc := 0
			count := int(math.Ceil(float64(l) / float64(settings.Threads)))
			for i := 0; i < l; i += count {
				n := i + count
				if n > l {
					n = l
				}
				// Println(q[i:n])
				m.A = m.A[i:n]
				go func(t1, t2 Matrix, calc int) {
					defer wg.Done()

					t := time.Now()
					m3 := t1.MulBy(&t2)
					Println(time.Since(t))
					results[calc] = Sprintf("%v\n", m3.PP())
				}(m.A, m.B, calc)
				calc++
			}

			wg.Wait()
		default:
			results = append(results, "SOMETHING WRONG")
		}

		ss, _ := json.Marshal(results)
		Fprintf(w, "%s", string(ss))
	})

	http.ListenAndServe(":1323", nil)
}
