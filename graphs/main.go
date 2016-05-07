package main

import (
	"crypto/md5"
	"fmt"
	"github.com/Jeffail/gabs"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var key string = "84c9a5c4-d59b-41d7-8b78-1794d43d3549"

// var key string = "aa6a7c24-03d4-4c7b-8b13-3813d4413663"

func Search(name string) Points {

	resp, _ := http.Get(fmt.Sprintf("https://search-maps.yandex.ru/v1/?apikey=%s&text=%s&lang=ru_RU", key, name))
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	jsonParsed, _ := gabs.ParseJSON(content)
	// fmt.Println(jsonParsed)

	ser := jsonParsed.S("features")
	k, _ := ser.Children()

	res := Points{}
	h := md5.New()

	for _, children := range k {
		h.Reset()
		name := children.Path("properties.CompanyMetaData.name").Data().(string)
		coords := children.Path("geometry.coordinates")
		io.WriteString(h, name)
		hash := h.Sum(nil)
		p := Point{}
		p.Y = coords.Index(0).Data().(float64)
		p.X = coords.Index(1).Data().(float64)
		p.Name = name
		p.Hash = fmt.Sprintf("%x", hash)
		res = append(res, p)
	}

	return res
}

func main() {
	finds := []string{
		"больницы кургана",
		"детские поликлиники кургана",
	}

	res := Points{}
	for _, find := range finds {
		res.Merge(Search(find))
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
	results := goPath(&paths, &res, 0)
	fmt.Printf("https://www.google.ru/maps/dir/%s", strings.Join(results, "/"))
}

func goPath(p *[]MiniPair, res *Points, first int) []string {
	var out []string
	temp := first
	for i := 0; i < len(*p); i++ {
		for j := 0; j < len(*p); j++ {
			if (*p)[j].x == temp {
				e := (*res).Find((*p)[j].x)
				out = append(out, fmt.Sprintf("%f+%f", e.X, e.Y))
				temp = (*p)[j].y
				break
			}
		}
	}
	out = append(out, out[0])
	return out
}

func sort(p *[]MiniPair) {
	for i := len(*p) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if (*p)[j].x > (*p)[j+1].x {
				x := MiniPair{(*p)[j].x, (*p)[j].y}
				(*p)[j] = (*p)[j+1]
				(*p)[j+1] = x
			}
		}
	}
}
