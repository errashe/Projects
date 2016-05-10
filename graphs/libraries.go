package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ehandle(where string, e error) {
	if e != nil {
		fmt.Println(where, " -> ", e.Error())
		os.Exit(0)
	}
}

func goPath(p *[]MiniPair, res *Points, first int) []string {
	var out []string
	s := first
	for i := 0; i < len(*p); i++ {
		e := (*res).Find((*p)[s].x)
		out = append(out, fmt.Sprintf("%f,%f", e.X, e.Y))
		s = (*p)[s].y
	}
	out = append(out, out[0])
	return out
}

// func goPath(p *[]MiniPair, res *Points, first int) []string {
// 	var out []string
// 	temp := first
// 	for i := 0; i < len(*p); i++ {
// 		for j := 0; j < len(*p); j++ {
// 			if (*p)[j].x == temp {
// 				e := (*res).Find((*p)[j].x)
// 				out = append(out, fmt.Sprintf("%f,%f", e.X, e.Y))
// 				temp = (*p)[j].y
// 				break
// 			}
// 		}
// 	}
// 	out = append(out, out[0])
// 	return out
// }

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

type slowed func()

func delaySecond(n time.Duration, fn slowed) {
	time.Sleep(n * time.Second)
	fn()
}

func start(what string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", what).Start()
	case "windows", "darwin":
		err = exec.Command("open", what).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	ehandle("OPEN", err)
}
