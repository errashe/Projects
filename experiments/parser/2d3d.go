package main

import "github.com/headzoo/surf"
import "os"
import "github.com/djimenez/iconv-go"
import "github.com/PuerkitoBio/goquery"
import "log"
import "fmt"
import "time"
import "math/rand"

func ruprint(in string) {
	str, _ := iconv.ConvertString(in, "windows-1251", "utf-8")
	println(str)
}

type fl struct {
	name string
	link string
}

func main() {

	lnks := make(chan string, 10000)
	attachs := make(chan fl, 50000)

	b := surf.NewBrowser()
	b.Open("http://www.2d-3d.ru/")

	form, _ := b.Form("div.login form")
	form.Input("login_name", "shade")
	form.Input("login_password", "Mm3363067")
	form.Submit()

	btn := b.Find("a.close")
	if btn.Text() != "×" {
		os.Exit(0)
	}

	for i := 1; i <= 126; i++ {
		b.Open(fmt.Sprintf("http://www.2d-3d.ru/2d-galereia/page/%d/", i))

		links := b.Find("div.shortstory p.lead a")
		links.Each(func(i int, e *goquery.Selection) {
			link, _ := e.Attr("href")
			lnks <- link
		})
		log.Println("Page", i, "parsed")
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
	close(lnks)

	for link := range lnks {
		b.Open(link)
		attach := b.Find("span.attachment a")
		if attach.Length() > 0 {
			href, _ := attach.Attr("href")
			attach := fl{attach.Text(), href}
			attachs <- attach
			// b.Click("span.attachment a")

			// filename := "files/" + attach.name
			// fout, err := os.Create(filename)
			// if err != nil {
			// 	log.Printf(
			// 		"Error creating file '%s'.", filename)
			// 	continue
			// }
			// defer fout.Close()

			// _, err = b.Download(fout)
			// if err != nil {
			// 	log.Printf(
			// 		"Error downloading file '%s'.", filename)
			// }
			// log.Println("Downloaded", attach.name)
		}
	}
	close(attachs)
	log.Println(len(attachs))
}
