package main

import (
	. "fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"github.com/headzoo/surf"
	"github.com/jessevdk/go-flags"
	"os"
	"strconv"
	"strings"
)

type Song struct {
	Link   string `csv:"link"`
	Artist string `csv:"artist"`
	Title  string `csv:"title"`
}

var opts struct {
	Email string `short:"e" long:"email" description:"Put email here" required:"true"`
	Pass  string `short:"p" long:"pass" description:"Put password here" required:"true"`
}

var err error

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(0)
	}

	bow := surf.NewBrowser()

	err = bow.Open("http://m.vk.com/")
	if err != nil {
		Println("Can't open login page: ", err)
		os.Exit(0)
	}

	login_form := bow.Forms()[1]

	login_form.Input("email", opts.Email)
	login_form.Input("pass", opts.Pass)
	err = login_form.Submit()
	if err != nil {
		Println("Can't login in: ", err)
		os.Exit(0)
	}

	err = bow.Open("https://m.vk.com/id0")
	if err != nil {
		Println("Can't open profile: ", err)
	}
	count, err := strconv.Atoi(bow.Find("ul.profile_menu li a[href*=audios] em").Text())
	if err != nil {
		Println("Can't find count: ", err)
		os.Exit(0)
	}

	export := []*Song{}

	for i := 0; i < count; i += 50 {
		err = bow.Open(Sprintf("http://m.vk.com/audio?offset=%d", i))
		if err != nil {
			Println("Can't open audio page: ", err)
			os.Exit(0)
		}

		songs := bow.Find("div.audio_item")

		songs.Each(func(i int, song *goquery.Selection) {
			link, exists := song.Find("input[type=hidden]").Attr("value")
			if !exists {
				Println("Can't find songs")
				os.Exit(0)
			}

			export = append(export, &Song{
				Link:   strings.Split(link, "?")[0],
				Artist: song.Find("span.ai_artist").Text(),
				Title:  song.Find("span.ai_title").Text(),
			})
		})
	}

	clientsFile, err := os.OpenFile("songs.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		Println("Can't open file: ", err)
	}
	defer clientsFile.Close()

	err = gocsv.Marshal(export, clientsFile)
	if err != nil {
		Println("Can't marshal export: ", err)
	}

	Printf("Done %d songs\n", len(export))
}
