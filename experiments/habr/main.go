package main

import (
	. "fmt"
	"net/http"

	"github.com/Jeffail/gabs"
)

func main() {
	res, _ := http.Get("https://tmfeed.ru/api/v1/habrahabr-geektimes_all_alltime.json")

	c, _ := gabs.ParseJSONBuffer(res.Body)

	m, _ := c.Path("posts.hubs").Children()

	for _, post := range m {
		hubs, _ := post.Children()
		for _, hub := range hubs {
			Println(hub.Path("title").Data().(string))
		}
		Println("--")
	}
}
