package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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
		t, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, []byte(session.Get("test").(string)))
	}
}

func main() {
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("test", "SOMEVALUEHERE")
		session.Save()
		c.HTML(200, "index.html", nil)
	})

	r.GET("/test", func(c *gin.Context) {
		session := sessions.Default(c)
		c.String(200, "%s - is a fuckin string", session.Get("test").(string))
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c)
	})

	r.Run(":8080")
}
