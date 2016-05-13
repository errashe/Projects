package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func ehandle(where string, e error) {
	if e != nil {
		fmt.Println(where, " -> ", e.Error())
	}
}

func delaySecond(n time.Duration, fn func()) {
	time.Sleep(n * time.Second)
	fn()
}

func start(what string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		// err = exec.Command("xdg-open", what).Start()
	case "windows":
		exec.Command("cmd", "/c", "start", what).Start()
	case "darwin":
		err = exec.Command("open", what).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	ehandle("OPEN", err)
}
