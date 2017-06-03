package main

import (
	. "fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	Println("Hello, world!")
	Println(rand.Intn(10))
}
