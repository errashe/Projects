package types

import (
	. "fmt"
	"sync"
)

type UserData struct {
	Server, Login, Password string
}

type Brute struct {
	Tasks chan UserData
}

func (b *Brute) Fill(data []UserData) {
	(*b).Tasks = make(chan UserData, 1e6)
	for _, d := range data {
		(*b).Tasks <- d
	}
	close(b.Tasks)
}

func (b Brute) Run(count int) {
	var wg sync.WaitGroup

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(b Brute) {
			defer wg.Done()
			for data := range b.Tasks {
				// TODO: WRITE HERE CHECK FOR SSH LOGIN/PASSWORD
				Println(data)
			}
		}(b)
	}

	wg.Wait()
}
