package main

import (
	"github.com/Jeffail/gabs"
	"net/http"
	"net/url"
)

func main() {
	j := gabs.New()

	j.SetP(2, "settings.cores")
	j.SetP(2, "settings.threads")

	j.SetP(3, "work")

	http.PostForm("http://localhost:1323/", url.Values{"payload": {j.String()}})
}
