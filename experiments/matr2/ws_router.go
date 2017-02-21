package main

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	. "net/http"
	. "strings"
)

func ws_router() *gin.Engine {
	r := gin.Default()

	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *Request) bool { return true }

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		Println(s.Request.Cookies())
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		msgStr := string(msg)

		commandStr := Split(msgStr, "|")

		if len(commandStr) <= 1 {
			s.Write([]byte("error|Message looks like a shit!"))
			return
		}

		switch commandStr[0] {
		case "test":
			s.Write([]byte("text|just text"))
		case "btest":
			m.Broadcast([]byte("text|just btext"))
		default:
			s.Write([]byte("error|function not found"))
		}
	})

	return r
}
