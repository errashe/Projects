package main

import (
	"fmt"

	"github.com/headzoo/surf"
)

func main() {
	a := surf.NewBrowser()

	a.Open("http://google.com/")
	form := a.Forms()[1]
	form.Input("q", "Hello, world!")
	form.Submit()

	fmt.Println(a.Title())
}
