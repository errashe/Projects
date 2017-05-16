package main

import . "fmt"
import "net/http"
import "time"
import "io/ioutil"
import "strings"

import "github.com/PuerkitoBio/goquery"

func main() {
	sites := []string{"avast.com"}

	client := http.Client{}
	client.Timeout = time.Second

	for _, site := range sites {
		res, err := client.Get(Sprintf("http://%s/", site))
		if err != nil {
			Println(err)
			return
		}

		bdy, _ := ioutil.ReadAll(res.Body)

		Println(res.Header)

		goq, err := goquery.NewDocumentFromReader(strings.NewReader(string(bdy)))

		var tit string = goq.Find("title").Text()

		Println(tit)
	}
}
