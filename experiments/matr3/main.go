package main

import (
	// . "fmt"
	"github.com/olahol/melody"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.LoadHTMLFiles("./main.html")
	r.StaticFS("/assets", http.Dir("./assets"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "main.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":8000")
}
