package main

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func downloadFromUrl(url, fileName string) {
	fmt.Println("Downloading", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println("Downloaded", fileName)
}

var token string = "18125f3dea048c0dfc3b75b906691ee20a603e56ba0cdaaefcdca43733602d1e96b07c7a199e3077a175f"

func VK(method, params string) []byte {
	resp, _ := http.Get(fmt.Sprintf("https://api.vk.com/method/%s?%s&access_token=%s", method, params, token))
	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)

	return str
}

func main() {
	jsonParsed, _ := gabs.ParseJSON(VK("audio.get", "owner_id=76822135&count=404"))

	ser := jsonParsed.S("response")
	k, _ := ser.Children()

	for _, song := range k[1:] {
		artist := song.Search("artist")
		title := song.Search("title")
		url := song.Search("url")
		downloadFromUrl(url.Data().(string), fmt.Sprintf("music/%s - %s.mp3", artist.Data(), title.Data()))
	}
}
