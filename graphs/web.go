package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Points []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("t")
	t, _ = t.Parse(string(mainHtml))
	t, _ = t.Parse(string(formHtml))
	t.Execute(w, "qwe")
}

func ghandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	res := StartSearch(r.Form["path"])
	if len(res) == 0 {
		fmt.Fprintf(w, "Не достаточно аргументов или не найдены объекты по критериям")
		return
	}
	t := template.New("t")
	t, _ = t.Parse(string(mainHtml))
	t, _ = t.Parse(string(graphHtml))

	t.Execute(w, &Page{res})
}

func StartServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/graph", ghandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
