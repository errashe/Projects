package main

import "github.com/headzoo/surf"

func main() {
	a := surf.NewBrowser()

	a.Open("http://yandex.ru/")

	frms := a.Forms()
	f := frms[len(frms)-1]

	f.Input("text", "Привет")
	f.Submit()
	for _, link := range a.Links() {
		println(link.Text)
	}
}
