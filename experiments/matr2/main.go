package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	m := main_router()
	w := ws_router()
	f := file_router()

	wg.Add(3)

	go m.Run(":8000")
	go w.Run(":8001")
	go f.Run(":8002")

	wg.Wait()
}
