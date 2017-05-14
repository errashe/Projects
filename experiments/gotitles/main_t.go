package main

import (
	. "fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Task struct {
	Url      string
	Attempts int
}

func (t Task) String() string {
	return t.Url
}

func (t Task) Full() string {
	return Sprintf("%s - %d", t.Url, t.Attempts)
}

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

var trigger int

var wg sync.WaitGroup
var results []string
var m sync.Mutex

func tch(val int) {
	m.Lock()
	trigger += val
	m.Unlock()
}

var (
	infile  = kingpin.Flag("infile", "Файл из которого следует читать").Short('i').Required().File()
	outfile = kingpin.Flag("outfile", "Файл в который следует писать (будет перезаписан)").Short('o').Required().String()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	t := time.Now()

	kingpin.Parse()

	file, _ := ioutil.ReadAll(*infile)
	lines := strings.Split(string(file), "\n")
	lines = removeDuplicates(lines)

	sort.Sort(sort.StringSlice(lines))

	tasks := make(chan Task, 1)
	defer close(tasks)

	go func(lines []string) {
		for _, line := range lines {
			tmpLine := strings.Split(line, "/")
			tasks <- Task{tmpLine[0], 0}
		}
	}(lines)

	threads := 2

	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			c := &http.Client{}
			c.Timeout = time.Second

			for {
				select {
				case task, ok := <-tasks:
					if !ok {
						return
					}
					if task.Attempts >= 5 {
						continue
					}

					Println(task.Full())

					res, err := c.Get(Sprintf("http://%s/", task.Url))
					if err != nil {
						Println(err)
						task.Attempts++
						tasks <- task
						continue
					}

					goq, err := goquery.NewDocumentFromResponse(res)
					if err != nil {
						Println(err)
						task.Attempts++
						tasks <- task
						continue
					}

					m.Lock()
					results = append(results, Sprintf("%s\t%s", task.Url, goq.Find("title").Text()))
					m.Unlock()
				default:
					continue
				}
			}
		}()
	}

	wg.Wait()

	Println(strings.Join(results, "\n"))

	ioutil.WriteFile(*outfile, []byte(strings.Join(results, "\n")), 0644)

	Println(time.Since(t))
}
