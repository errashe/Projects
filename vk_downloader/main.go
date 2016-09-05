package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

func he(e error) {
	if e != nil {
		println(e.Error())
	}
}

func main() {
	a := surf.NewBrowser()

	e := a.Open("https://m.vk.com")
	he(e)

	form := a.Forms()[1]
	form.Input("email", "e4stw00d@icloud.com")
	form.Input("pass", "the{}Pre4cher")
	form.Submit()

	res_count := 0

	work := true
	i := 0
	for work {
		e = a.Open(fmt.Sprintf("https://m.vk.com//audio?offset=%d", i*50))
		he(e)
		i++

		audios := a.Find("div.audio_item")

		if audios.Length() == 0 {
			work = !work
			break
		}

		audios.Each(func(i int, audio *goquery.Selection) {
			el := audio.Find("div.ai_body input[type=hidden]").First()
			_, e := el.Attr("value")

			if e {
				res_count++
			}
		})

	}
	println(res_count)

}
