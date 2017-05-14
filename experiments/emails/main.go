package main

import . "fmt"
import "github.com/simia-tech/go-pop3"
import "github.com/veqryn/go-email/email"
import "crypto/tls"
import "strings"

var login = "defensuer@gmail.com"
var password = "Open!3451"

func main() {
	c, err := pop3.Dial("pop.gmail.com:995", pop3.UseTLS(&tls.Config{}))
	if err != nil {
		Println(err)
	}

	defer c.Quit()

	err = c.Auth(login, password)
	if err != nil {
		Println(err)
	}

	Println(c.Stat())

	msgs, _ := c.ListAll()
	Println(len(msgs), msgs)
	text, err := c.Retr(msgs[0].Seq)
	if err != nil {
		Println(err)
	}

	msg, err := email.ParseMessage(strings.NewReader(text))
	if err != nil {
		Println(err)
	}
	Println(msg.MessagesAll()[0])

	// msgs, err := c.ListAll()
	// if err != nil {
	// 	Println(err)
	// }

	// for _, msg := range msgs {
	// 	str, _ := c.Retr(msg.Seq)
	// 	Println(str)
	// 	Println("\n\n-------------------------------------------------------------------------------------\n\n")
	// }
}
