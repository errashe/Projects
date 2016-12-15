package main

import "fmt"
import . "./database"

func main() {
	fmt.Println("Hello, world!")
	Something()
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	fmt.Println(i, j)
}
