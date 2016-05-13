package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hmain(w http.ResponseWriter, r *http.Request) {
	t := template.New("tmpl")
	t, _ = t.Parse(string(mainHtml))
	t.Execute(w, "world")
}

func main() {
	http.HandleFunc("/", hmain)
	delaySecond(1, func() {
		fmt.Println(":8080")
		start("http://localhost:8080/")
	})
	http.ListenAndServe(":8080", nil)
}
