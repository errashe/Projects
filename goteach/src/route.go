package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main_handler(c *gin.Context) {
	session, err := store.Get(c.Request, "main")
	error_handler(err)

	if session.Values["username"] == nil {
		f := session.Flashes()
		session.Save(c.Request, c.Writer)
		c.HTML(200, "data/main.html", gin.H{"f": f})
	} else {
		c.HTML(200, "data/chat.html", nil)
	}

}

func login_handler(c *gin.Context) {
	session, err := store.Get(c.Request, "main")
	error_handler(err)

	if checkUser(c.PostForm("username"), c.PostForm("password")) {
		session.Values["username"] = c.PostForm("username")
	} else {
		session.AddFlash("FUCKIN ERROR")
	}

	session.Save(c.Request, c.Writer)
	c.Redirect(302, "/")
}

func logout_handler(c *gin.Context) {
	session, err := store.Get(c.Request, "main")
	error_handler(err)

	session.Values["username"] = nil
	session.Save(c.Request, c.Writer)

	c.Redirect(302, "/")
}

func ws_handler(c *gin.Context) {
	session, err := store.Get(c.Request, "main")
	error_handler(err)

	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	error_handler(err)

	conns = append(conns, conn)

	getLastFifteenMessages(conn)

	for {
		_, r, err := conn.NextReader()
		if err != nil {
			break
		}

		msg := make([]byte, 1024)

		n, err := r.Read(msg)
		error_handler(err)

		addMessage(session.Values["username"].(string), string(msg[:n]))
	}
	conn.Close()
}
