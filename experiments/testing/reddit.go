package main

import "fmt"
import "sync"
import "time"
import "github.com/headzoo/surf"

func main() {
	var wg = &sync.WaitGroup{}
	urls := []string{"http://google.com/", "http://yandex.ru/", "http://wowjp.net/", "http://tmfeed.ru/"}
	var ch_urls = make(chan string, len(urls))
	for _, u := range urls {
		ch_urls <- u
	}
	close(ch_urls)

	var th_cnt int = 10
	b := surf.NewBrowser()

	t := time.Now()
	for i := 0; i < th_cnt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range ch_urls {
				b.Open(url)
				fmt.Printf("%s\n", b.Title())
			}
		}()
	}

	wg.Wait()
	fmt.Printf("%s\n", time.Since(t))
}
