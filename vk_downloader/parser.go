package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	r "gopkg.in/dancannon/gorethink.v2"
)

func he(e error) {
	if e != nil {
		println(e.Error())
	}
}

func main() {
	a := surf.NewBrowser()
	s, _ := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})

	table := r.DB("experiments").Table("audios")

	a.Open("https://m.vk.com")

	form := a.Forms()[1]
	form.Input("email", "e4stw00d@icloud.com")
	form.Input("pass", "the{}Pre4cher")
	form.Submit()

	work := true
	i := 0
	for work {
		a.Open(fmt.Sprintf("https://m.vk.com//audio?offset=%d", i*50))
		i++

		audios := a.Find("div.audio_item")

		if audios.Length() == 0 {
			work = !work
			break
		}

		audios.Each(func(i int, audio *goquery.Selection) {
			body := audio.Find("div.ai_body").First()
			el := body.Find("input[type=hidden]").First()
			url, e := el.Attr("value")

			label := body.Find("div.ai_label").First()

			el = label.Find("span.ai_artist").First()
			artist := el.Text()
			el = label.Find("span.ai_title").First()
			title := el.Text()

			if e {
				table.Insert(map[string]string{
					"url":    url,
					"artist": artist,
					"title":  title,
				}).Exec(s)
			}
		})

	}

}
