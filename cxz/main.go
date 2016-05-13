package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	// "strings"
)

func main() {
	// var lst []string

	a := surf.NewBrowser()
	a.Open("http://cxz.to/serials/fl_action_foreign_russion_hight_studio/")
	posters := a.Find("div.b-poster-tile")
	posters.Each(func(_ int, s *goquery.Selection) {
		l := s.Find(".b-poster-tile__link")
		link, _ := l.Attr("href")
		// lnk := strings.Split(link, "/")
		// llnk := strings.Split(lnk[2], "-")

		a.Open(fmt.Sprintf("http://cxz.to%s", link))
		container := a.Find("div.b-files-folders")
		fmt.Println(container.Html())

		// movies := container.Find(fmt.Sprintf("a[href*='%s']", llnk[0]))
		// movies.Each(func(_ int, s *goquery.Selection) {
		// 	fmt.Println(s.Attr("href"))
		// })
	})
}
