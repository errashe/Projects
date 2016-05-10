package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"io/ioutil"
	"os"
	"strings"
)

func read_map(map_str string) (*MapData, int, int) {
	rows := strings.Split(map_str, "\n")
	if len(rows) == 0 {
		panic("The map needs to have at least 1 row")
	}
	row_count := len(rows)
	col_count := len(rows[0])

	result := *NewMapData(row_count, col_count)
	for i := 0; i < row_count; i++ {
		for j := 0; j < col_count; j++ {
			char := rows[i][j]
			switch char {
			case 'L':
				result[i][j] = LAND
			case 'W':
				result[i][j] = WALL
			case 'S':
				result[i][j] = START
			case 'E':
				result[i][j] = STOP
			}
		}
	}
	return &result, row_count, col_count
}

func str_map(data *MapData, nodes []*Node, rows, cols int) [][]string {
	var result [][]string = make([][]string, rows)
	for i, row := range *data {
		result[i] = make([]string, cols)
		for j, cell := range row {
			switch cell {

			case LAND:
				result[i][j] = "L"
			case WALL:
				result[i][j] = "W"
			case START:
				result[i][j] = "S"
			case STOP:
				result[i][j] = "E"
			default:
				result[i][j] = "?"

			}
		}
	}

	if len(nodes) == 0 {
		fmt.Println("No path")
		os.Exit(0)
	}

	for _, node := range nodes[1 : len(nodes)-1] {
		result[node.x][node.y] = "P"
	}

	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough params")
		os.Exit(0)
	}

	filepath := os.Args[1]
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(0)
	}

	byt, _ := ioutil.ReadFile(os.Args[1])
	str := string(byt)
	m, rows, cols := read_map(str)
	g := NewGraph(m)
	n := Astar(g)

	dc := gg.NewContext(rows*40+40, cols*40+40)

	n_m := str_map(m, n, rows, cols)
	for x, row := range n_m {
		fmt.Println(row)
		for y, cell := range row {
			dc.Push()
			switch cell {
			case "L":
				dc.SetRGBA(.5, 1, .5, 1)
			case "W":
				dc.SetRGBA(1, 1, 1, 1)
			case "S":
				dc.SetRGBA(1, 0, 0, 1)
			case "E":
				dc.SetRGBA(0, 0, 1, 1)
			case "P":
				dc.SetRGBA(.5, .5, .5, 1)
			}
			dc.DrawRectangle(10+float64(y)*(30+10), 10+float64(x)*(30+10), 30, 30)
			dc.Fill()
			dc.Pop()
		}
	}

	dc.SavePNG("out.png")

}
