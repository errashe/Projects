package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max float64) float64 {
	return rand.Float64()*(max-min+1) + min
}

func GenerateMap(width, height int) [][]int {
	rand.Seed(time.Now().Unix())

	var res [][]int = make([][]int, width)
	for i := range res {
		res[i] = make([]int, height)
	}

	for x := 0; x < len(res); x++ {
		for y := 0; y < len(res[x]); y++ {
			if random(0, 100) < 30 {
				res[x][y] = 1
			} else {
				res[x][y] = 0
			}
		}
	}

	return res
}

func saveMap(m [][]int) [][]string {
	var res [][]string = make([][]string, len(m))
	for i := range res {
		res[i] = make([]string, len(m[i]))
	}

	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			switch m[x][y] {
			case 0:
				res[x][y] = "L"
			case 1:
				res[x][y] = "W"
			}
		}
	}

	return res
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Не определено имя файла")
		os.Exit(0)
	}
	os.Remove(os.Args[1])
	f, _ := os.Create(os.Args[1])

	var width int
	var height int
	if len(os.Args) >= 3 {
		width, _ = strconv.Atoi(os.Args[2])
		height, _ = strconv.Atoi(os.Args[3])
	} else {
		width = 10
		height = 10
	}

	w := bufio.NewWriter(f)
	m := saveMap(GenerateMap(width, height))

	var temp []string = make([]string, len(m))
	for i, row := range m {
		temp[i] = strings.Join(row, "")
	}
	res := strings.Join(temp, "\n")
	w.WriteString(res)
	w.Flush()
}
