package main

import (
	"fmt"
)

func main() {
	fmt.Println("Init program...")

	net := newNetwork()
	net.think([]float64{1, 1})

	fmt.Printf("%+v\n", net)
}
