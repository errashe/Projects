package main

import "log"
import "sync"
import "syscall"
import "time"

var wg sync.WaitGroup
var cnt int

func timer() {
	defer wg.Done()
	wg.Add(1)
	for x := range time.Tick(time.Second) {
		log.Println(cnt, x)
		cnt = 0
	}
}

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("Socket: ", err.Error())
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		log.Fatal("SetsockoutInt: ", err.Error())
	}

	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_NODELAY, 1)
	if err != nil {
		log.Fatal("SetsockoutInt: ", err.Error())
	}

	addr := syscall.SockaddrInet4{
		Port: 1234,
		Addr: [4]byte{127, 0, 0, 1},
	}

	err = syscall.Bind(fd, &addr)
	if err != nil {
		log.Fatal("Bind: ", err.Error())
	}

	err = syscall.Listen(fd, 5)
	if err != nil {
		log.Fatal("Listen: ", err.Error())
	}

	go timer()
	work := make(chan int, 1000000)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			var nfd int
			nfd, _, err = syscall.Accept(fd)
			if err != nil {
				log.Fatal("Accept: ", err.Error())
			}

			work <- nfd
		}
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for chfd := range work {
					b := make([]byte, 1024)
					_, _, err := syscall.Recvfrom(chfd, b, 0)
					if err != nil {
						log.Fatal("Recvfrom: ", err.Error())
					}

					// log.Println(string(b[:m]))
					cnt++
					syscall.Close(chfd)
				}
			}()
		}
	}()

	wg.Wait()
	syscall.Close(fd)
}
