package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(c *gin.Context) {
	session := sessions.Default(c)

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, []byte(fmt.Sprintf("%s:%s", session.Get("nick"), msg)))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("nick") == nil {
			session.Set("nick", RandStringRunes(8))
			session.Save()
		}
		c.HTML(200, "index.html", nil)
	})

	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)

		session.Delete("nick")
		session.Save()
		c.Redirect(302, "/")
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c)
	})

	r.Run(":8080")
}
