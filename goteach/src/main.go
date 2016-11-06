package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

var (
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	conns []*websocket.Conn
)

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(loadTemplates())

	initDB()
	defer db.Close()

	r.GET("/", main_handler)
	r.POST("/login", login_handler)
	r.GET("/logout", logout_handler)
	r.GET("/ws", ws_handler)

	go control()

	r.Run()
}
