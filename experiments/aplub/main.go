package main

import "os"
import "strconv"

func main() {
	i := 0
	for _, arg := range os.Args[1:] {
		temp, _ := strconv.Atoi(arg)
		i += temp
	}
	println(i)
}
