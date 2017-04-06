package main

import . "fmt"
import "net"
import "time"
import "sync"
import "io/ioutil"
import "strings"

import "gopkg.in/alecthomas/kingpin.v2"

var (
	file    = kingpin.Arg("server_file", "file with servers").Required().File()
	timeout = kingpin.Arg("timeout", "Time to check per port").Default("1").Int()
)

// import "os"

var num_threads int = 50

// var t_ports []int = []int{21, 22, 23, 25, 80, 110, 3306, 27017}
var t_ports []int = []int{22}
var wg sync.WaitGroup

func main() {
	t := time.Now()

	kingpin.Parse()

	file_b, _ := ioutil.ReadAll(*file)
	file_s := strings.TrimSpace(string(file_b))
	t_servers := strings.Split(file_s, "\n")

	jobs := make(chan string, len(t_servers)*len(t_ports))

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(jobs)
		for _, server := range t_servers {
			for _, port := range t_ports {
				jobs <- Sprintf("%s:%d", server, port)
			}
		}
	}()

	wg.Wait()

	wg.Add(num_threads)

	for i := 0; i < num_threads; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				n, e := net.DialTimeout("tcp4", job, time.Duration(*timeout)*time.Second)
				if e != nil {
					// Printf("OFFLINE %s\n", job)
				} else {
					Printf("%s\n", job)
					n.Close()
				}
			}
		}()
	}

	wg.Wait()

	Println(time.Since(t))
	Println(float64(len(t_servers)*len(t_ports)) / time.Since(t).Seconds())
}
