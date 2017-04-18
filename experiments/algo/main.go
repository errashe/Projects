package main

import (
	. "./types"
	. "fmt"
	"time"
)

func main() {
	// m1 := NewMatrix(3, 3, true)
	// m2 := NewMatrix(3, 3, true)
	// m3 := m1.MulByMatrix(m2)
	// v1 := NewVVector(3, true)
	// v2 := m1.MulByVVector(v1)
	// v3 := NewHVector(3, true)
	// v4 := v1.MulByHVector(v3)

	// Println(m1.String())
	// Println("#####")
	// Println(m2.String())
	// Println("#####")
	// Println(m3.String())
	// Println("#####")
	// Println("v1", v1.String())
	// Println("#####")
	// Println("v2", v2.String())
	// Println("#####")
	// Println("v3", v3.String())
	// Println("#####")
	// Println(v4.String())
	// Println("#####\n#####")

	// b := Brute{}
	// b.Fill([]UserData{
	// 	UserData{"localhost:22", "login", "password1"},
	// 	UserData{"localhost:22", "login", "password2"},
	// 	UserData{"localhost:22", "login", "password3"},
	// 	UserData{"localhost:22", "login", "password4"},
	// })
	// b.Run(2)

	m1 := NewMatrix(500, 1000, true)
	m2 := NewMatrix(1000, 1000, true)

	t := time.Now()
	m1.MulByMatrix(m2)
	Println(time.Since(t))
}
