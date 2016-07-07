package main

import "os"

func ehandle(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
}
