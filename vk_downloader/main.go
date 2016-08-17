package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Resp struct {
	Response []interface{} `json:"response"`
}

type Audio struct {
	Aid      int    `json:"aid"`
	Artist   string `json:"artist"`
	Duration int    `json:"duration"`
	Genre    int    `json:"genre"`
	OwnerID  int    `json:"owner_id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
}

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

func VK(method, params string) *Resp {
	url := fmt.Sprintf("https://api.vk.com/method/%s?%s&access_token=%s", method, params, token)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)

	ret := &Resp{}
	json.Unmarshal(str, ret)
	return ret
}

func main() {
	resp := VK("audio.get", "owner_id=76822135&count=1000")
	// count := resp.Response[0].(float64)
	audios := resp.Response[1:10].([]Audio)

	fmt.Println(audios)

}
