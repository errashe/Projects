package main

import "fmt"
import "sync"
import "os"
import "time"
import "strconv"

func main() {
	var i int = 0
	var wg sync.WaitGroup
	var m sync.Mutex

	t := time.Now()

	wg.Add(1)
	go func() {
		for {
			if i > 5000 {
				fmt.Println(time.Since(t))
				os.Exit(0)
			}
		}
	}()

	end, _ := strconv.ParseInt(os.Args[1], 10, 64)
	for a := 0; a < int(end); a++ {
		wg.Add(1)
		go func() {
			for {
				m.Lock()
				i++
				m.Unlock()
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	wg.Wait()
}
