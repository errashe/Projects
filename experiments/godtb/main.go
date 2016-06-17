package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

func main() {
	a := surf.NewBrowser()
	a.Open("http://www.dotabuff.com/players/92413647")
	rows := a.Find("div.r-table div.r-row[data-link-to*='/matches/']")
	rows.Each(func(i int, g *goquery.Selection) {
		link, isset := g.Attr("data-link-to")
		if isset {
			fmt.Println(link)
		}
	})
}
