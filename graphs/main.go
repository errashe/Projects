package main

import (
	"crypto/md5"
	"fmt"
	"github.com/Jeffail/gabs"
	"io"
	"io/ioutil"
	// "math"
	"net/http"
)

// var key string = "84c9a5c4-d59b-41d7-8b78-1794d43d3549"
var key string = "aa6a7c24-03d4-4c7b-8b13-3813d4413663"

func getByName(name string) Points {

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
		"музеи кургана",
		"достопримечательности кургана",
		"церкви кургана",
	}
	res := Points{}

	for _, find := range finds {
		res.Merge(getByName(find))
	}

	res.Numerate()

	count := len(res)

	pairs := Pairs{}

	for i := range res {
		for j := range res {
			p := Pair{}
			p.first = res[i]
			p.second = res[j]
			// if !pairs.Isset(p.Reverse()) && p.first != p.second {
			p.CalcDist()
			pairs = append(pairs, p)
			// }
		}
	}

	// for _, point := range res {
	// 	fmt.Println(point.Id+1, point.X, point.Y, point.Name)
	// }

	matrix := Matrix{}
	matrix.FillEmptyMatrix(count)
	matrix.FillByPairs(pairs)

	temp := Matrix{}
	temp.FillEmptyMatrix(count)
	temp.FillByPairs(pairs)

	var paths []MiniPair

	for x := 0; x < count-1; x++ {
		matrix.ReduxRows()
		matrix.ReduxCols()
		matrix.CalcKoef()
		mpairs := matrix.FindMaxKoef()

		max, iq, jq := 0, 0, 0
		for _, mpair := range mpairs {
			if temp[mpair.x][mpair.y].dist >= max {
				max = temp[mpair.x][mpair.y].dist
				iq = mpair.x
				jq = mpair.y
			}
		}
		paths = append(paths, MiniPair{iq, jq})
		matrix.DeleteRowCol(OldToNew(&matrix, iq, jq))
	}
	paths = append(paths, MiniPair{matrix[0][0].pos.x, matrix[0][0].pos.y})

	sort(&paths)
	goPath(&paths, &res, 0)
}

func goPath(p *[]MiniPair, res *Points, first int) {
	temp := first
	for i := 0; i < len(*p); i++ {
		for j := 0; j < len(*p); j++ {
			if (*p)[j].x == temp {
				// fmt.Println((*p)[j].x, "->", (*p)[j].y)
				r1, r2 := (*res).Find((*p)[j].x)
				fmt.Printf("%f+%f/", r1, r2)
				temp = (*p)[j].y
				break
			}
		}
	}
	fmt.Printf("\n")
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

func OldToNew(m *Matrix, iq, jq int) (int, int) {
	for a, row := range *m {
		for b, cell := range row {
			if cell.pos.x == iq && cell.pos.y == jq {
				return a, b
			}
		}
	}
	return 0, 0
}
