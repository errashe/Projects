package main

import (
	. "fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	wg sync.WaitGroup
	m  sync.Mutex
)

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

type Result struct {
	Title string
	Url   string
}

type Results []Result

func (r *Results) toString() string {
	var tmp []string
	for _, result := range *r {
		tmp = append(tmp, Sprintf("%s\t%s", result.Title, result.Url))
	}
	return strings.Join(tmp, "\n")
}

func (r *Result) String() string {
	return Sprintf("%s\t%s", r.Title, r.Url)
}

type Task struct {
	Url     string
	Attemts int
}

type Worker struct {
	Client  http.Client
	Tasks   []Task
	Results Results
}

func newWorker() Worker {
	c := http.Client{}
	c.Timeout = 5 * time.Second
	return Worker{c, []Task{}, Results{}}
}

func (w *Worker) Run(urls []string) {
	w.Fill(urls)

	wg.Add(1)
	go w.Do()
}

func (w *Worker) Fill(urls []string) {
	w.Tasks = []Task{}
	for _, url := range urls {
		if url == "" {
			continue
		}
		w.Tasks = append(w.Tasks, Task{url, 0})
	}
}

func (w *Worker) Do() {
	defer wg.Done()

	var task Task
	for len(w.Tasks) > 0 {
		task, w.Tasks = w.Tasks[0], w.Tasks[1:]

		if task.Attemts >= 5 {
			// Println("DELETE ", task.Url)
			continue
		}

		tout := time.Now()
		res, err := w.Client.Get(Sprintf("http://%s", task.Url))
		if err != nil || time.Since(tout) >= 5000*time.Millisecond {
			task.Attemts++
			w.Tasks = append(w.Tasks, task)
			// Println("RETRY ", task.Url)
			continue
		}

		bdy, _ := ioutil.ReadAll(res.Body)

		goq, err := goquery.NewDocumentFromReader(strings.NewReader(string(bdy))) // KOSTYL
		if err != nil {
			Println("GOQ", err)
			continue
		}
		w.Results = append(w.Results, Result{strings.TrimSpace(goq.Find("title").Text()), task.Url})
		// Println("APPEND ", task.Url)
	}
}

var (
	in      = kingpin.Flag("in", "in file").Short('i').Default("in.csv").String()
	out     = kingpin.Flag("out", "out file").Short('o').Default("out.csv").String()
	ru      = kingpin.Flag("ru", "need ru sorting/splitting").Short('r').Default("false").Bool()
	threads = kingpin.Flag("threads", "number of goroutines").Short('t').Default("10").Int()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	kingpin.Parse()

	log.SetOutput(ioutil.Discard)

	t := time.Now()

	fileB, err := ioutil.ReadFile(*in)
	if err != nil {
		Println(err)
		return
	}
	fileL := strings.Split(string(fileB), "\n")
	fileD := removeDuplicates(fileL)

	sort.Sort(sort.StringSlice(fileD))

	l := len(fileD)
	step := int(math.Ceil(float64(l) / float64(*threads)))

	var workers []*Worker

	for i := 0; i < l; i += step {
		tmp := newWorker()
		workers = append(workers, &tmp)
		workers[len(workers)-1].Run(fileD[i : i+step])
	}

	wg.Wait()

	var results Results
	for _, worker := range workers {
		results = append(results, worker.Results...)
	}
	ioutil.WriteFile("out.csv", []byte(results.toString()), 0644)
	Println("COMMON DONE")

	if *ru {

		var ru, en Results
		for _, result := range results {
			if result.Url[len(result.Url)-2:] == "ru" {
				ru = append(ru, result)
			} else {
				en = append(en, result)
			}
		}

		ioutil.WriteFile("ru.csv", []byte(ru.toString()), 0644)
		ioutil.WriteFile("en.csv", []byte(en.toString()), 0644)

		Println("RU/EN DONE")
	}

	Println("NOW U CAN CHECK out.csv AND ru.csv/en.csv")

	Println(time.Since(t))
}
