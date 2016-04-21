package main

import "fmt"
import "sync"
import "time"
import "math/rand"

type Philo struct {
	name string
}

var wg sync.WaitGroup
var table []sync.Mutex

func (p *Philo) Take(s int) {
	defer fmt.Println(p.name, "взял", s)
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	table[s].Lock()
}

func (p *Philo) Drop(s int) {
	defer fmt.Println(p.name, "освободил", s)
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	table[s].Unlock()
}

func (p *Philo) Eat() {
	defer fmt.Println(p.name, "ест")
	time.Sleep(200 * time.Millisecond)
}

func main() {

	table = []sync.Mutex{
		sync.Mutex{},
		sync.Mutex{},
	}

	philos := []Philo{
		Philo{"Первый"},
		Philo{"Второй"},
		Philo{"Третий"},
	}

	for _, p := range philos {
		wg.Add(1)
		go func(phil Philo) {
			defer wg.Done()
			for {
				phil.Take(0)
				phil.Take(1)
				phil.Eat()
				phil.Drop(0)
				phil.Drop(1)
			}
		}(p)
	}

	wg.Wait()
}
