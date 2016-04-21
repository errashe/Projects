package main

import "log"
import "strings"
import "os"
import "sync"
import "syscall"
import "strconv"

var wg sync.WaitGroup

func main() {

	if len(os.Args) < 2 {
		log.Fatal("FUCK")
	}
	straddr := strings.Split(os.Args[1], ".")
	var ad [4]byte
	for ch := range straddr {
		i, _ := strconv.Atoi(straddr[ch])
		ad[ch] = byte(i)
	}

	for a := 0; a < 100; a++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
				if err != nil {
					log.Fatal("Socket: ", err.Error())
				}
				addr := syscall.SockaddrInet4{
					Port: 1234,
					Addr: ad,
				}

				err = syscall.Connect(fd, &addr)
				if err != nil {
					log.Fatal("Connect: ", err.Error())
				}

				err = syscall.Sendto(fd, []byte("Hello"), 0, &addr)
				if err != nil {
					log.Fatal("Sendto: ", err.Error())
				}
				syscall.Close(fd)
			}
		}()
	}

	wg.Wait()
}
