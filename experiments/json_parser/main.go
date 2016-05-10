package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var key string = "84c9a5c4-d59b-41d7-8b78-1794d43d3549"

func ehandle(where string, err error) {
	if err != nil {
		fmt.Println(where, ": ", err.Error())
	}
}

func Search(name string) (string, []byte) {

	link := fmt.Sprintf("https://search-maps.yandex.ru/v1/?apikey=%s&text=%s&lang=ru_RU", key, name)
	resp, err := http.Get(link)
	ehandle("Get", err)
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	ehandle("ReadAll", err)

	return link, content
}

func main() {
	var p map[string]interface{}

	_, data := Search("музеи кургана")

	err := json.Unmarshal(data, &p)
	ehandle("JSON", err)

	res := p["features"].([]interface{})

	for _, item := range res {
		pname := item.(map[string]interface{})["properties"]
		cname := pname.(map[string]interface{})["CompanyMetaData"]
		name := cname.(map[string]interface{})["name"].(string)
		fmt.Println(name)
		geometry := item.(map[string]interface{})["geometry"]
		xy := geometry.(map[string]interface{})["coordinates"]
		x := xy.([]interface{})[1].(float64)
		y := xy.([]interface{})[0].(float64)

		fmt.Printf("X=%f, Y=%f\n", x, y)
	}

}
