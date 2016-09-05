package main

import (
	"fmt"
	r "gopkg.in/dancannon/gorethink.v2"
	"io"
	"net/http"
	"os"
)

var dir string = "music"

func main() {
	s, _ := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})

	var audios []map[string]string

	cur, _ := r.DB("experiments").Table("audios").Run(s)
	cur.All(&audios)

	for _, song := range audios {
		download(fmt.Sprintf("%s - %s.mp3", song["artist"], song["title"]), song["url"])
		fmt.Printf("%s - %s.mp3\n", song["artist"], song["title"])
	}

}

func download(filename, url string) {
	out, err := os.Create(fmt.Sprintf("%s/%s", dir, filename))
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
