package main

import . "fmt"
import "os"

import "gopkg.in/headzoo/surf.v1"
import "github.com/PuerkitoBio/goquery"

func main() {
	f, err := os.OpenFile("songs.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bow := surf.NewBrowser()
	bow.Open("http://m.vk.com/")

	login_form := bow.Forms()[1]
	login_form.Input("email", "e4stw00d@icloud.com")
	login_form.Input("pass", "Open!3451")
	login_form.Submit()
	if len(bow.Forms()) > 1 {
		Println("NEED MORE AUTH")
		login_form := bow.Forms()[1]
		login_form.Input("email", "e4stw00d@icloud.com")
		login_form.Input("pass", "Open!3451")
		sid, _ := login_form.Dom().Find("input[name=captcha_sid]").Attr("value")

		f, _ := os.OpenFile("pic.jpg", os.O_CREATE|os.O_WRONLY, 0644)
		bow.Open(Sprintf("http://m.vk.com/captcha.php?sid=%s", sid))
		bow.Download(f)
		f.Close()

		var usid string
		Print("INSERT CAPTCHA: ")
		Scan(&usid)

		login_form.Input("captcha_key", usid)
		err = login_form.Submit()
	}

	offset_len, offset := 50, 0

	for offset_len != 0 {
		bow.Open(Sprintf("https://m.vk.com/audio?offset=%d", offset))

		audio_items := bow.Find("div.audio_item")
		offset_len = audio_items.Length()
		audio_items.Each(func(i int, s *goquery.Selection) {
			artist := s.Find("span.ai_artist").Text()
			title := s.Find("span.ai_title").Text()
			link, _ := s.Find("input[type=hidden]").Attr("value")
			if link != "" {
				Fprintf(f, "%s|%s - %s.mp3\n", link, artist, title)
			}
		})

		offset += 50
	}
}
