package main

import (
	. "fmt"
	"github.com/unrolled/render"
	"html/template"
)

var (
	re *rnd.Render
)

func init() {
	Println("render loaded")

	re = rnd.New(rnd.Options{
		Layout:        "layout",
		Extensions:    []string{".tmpl"},
		Funcs:         []template.FuncMap{funcs},
		IsDevelopment: true,
	})
}
