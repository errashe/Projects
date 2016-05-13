package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var key string = "84c9a5c4-d59b-41d7-8b78-1794d43d3549"

func Search(name string) Points {

	resp, err := http.Get(fmt.Sprintf("https://search-maps.yandex.ru/v1/?apikey=%s&text=%s&lang=ru_RU&results=5000", key, name))
	ehandle("Get", err)
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	ehandle("ReadAll", err)

	var jsonParsed map[string]interface{}

	err = json.Unmarshal(content, &jsonParsed)
	ehandle("JSON", err)

	ser := jsonParsed["features"].([]interface{})

	res := Points{}
	h := md5.New()

	for _, item := range ser {
		h.Reset()

		pname := item.(map[string]interface{})["properties"]
		cname := pname.(map[string]interface{})["CompanyMetaData"]
		name := cname.(map[string]interface{})["name"].(string)
		geometry := item.(map[string]interface{})["geometry"]
		xy := geometry.(map[string]interface{})["coordinates"]

		io.WriteString(h, name)
		hash := h.Sum(nil)
		p := Point{}
		p.X = xy.([]interface{})[1].(float64)
		p.Y = xy.([]interface{})[0].(float64)
		p.Name = name
		p.Hash = fmt.Sprintf("%x", hash)
		res = append(res, p)
	}

	return res
}

func StartSearch(searchs []string) []string {
	if len(searchs) == 0 {
		return []string{}
	}

	res := Points{}
	for _, find := range searchs {
		res.Merge(Search(find))
	}

	if len(res) == 0 {
		return []string{}
	}

	res.Numerate()

	count := len(res)
	pairs := Pairs{}

	pairs.FillByPoints(&res)

	matrix := Matrix{}
	matrix.FillEmptyMatrix(count)
	matrix.FillByPairs(pairs)
	temp := matrix.FillTemp()

	var paths []MiniPair

	for x := 0; x < count-1; x++ {
		matrix.ReduxRows()
		matrix.ReduxCols()
		matrix.CalcKoef()
		iq, jq := matrix.FindMaxKoef(&temp)
		paths = append(paths, MiniPair{iq, jq})
		matrix.DeleteRowCol(matrix.OldToNew(iq, jq))
	}
	paths = append(paths, MiniPair{matrix[0][0].pos.x, matrix[0][0].pos.y})

	sort(&paths)
	return goPath(&paths, &res, 0)
}

func main() {
	go delaySecond(1, func() {
		start("http://localhost:8080/")
		fmt.Println("Served on 8080")
	})
	StartServer()
}
