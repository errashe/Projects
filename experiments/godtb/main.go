package main

import (
	"flag"
	"fmt"
	"github.com/headzoo/surf"
	"os"
)

func main() {
	usage := flag.Bool("h", false, "Показать помощь")
	link := flag.String("link", "http://www.dotabuff.com/players/92413647", "Ссылка на dotabuff пользователя")
	flag.Parse()

	if *usage {
		flag.Usage()
		os.Exit(0)
	}

	a := surf.NewBrowser()
	a.Open(*link)
	rows := a.Find("div.r-table div.r-row[data-link-to*='/matches/'] a[href*='/matches/']")
	cnt := 0.0
	for i := 0; i < rows.Length(); i += 2 {
		name := rows.Eq(i).Text()
		status := rows.Eq(i + 1).Text()
		symb := ""
		if status == "Won Match" {
			symb = "*"
			cnt++
		}
		fmt.Printf("%s%s - %s\n", symb, name, status)
	}
	fmt.Printf("%.2f%% win rate par last 15 matches!\n", cnt*100/15.0)
}
