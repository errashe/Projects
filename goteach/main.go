package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/gorilla/websocket"
	"html/template"
	"math/rand"
	"time"
)

var (
	conns []*websocket.Conn
)

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func index(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("nick") == nil {
		session.Set("nick", RandStringRunes(8))
		session.Save()
	}

	rows, err := AllMessages()
	if err != nil {
		fmt.Println(err.Error())
	}

	c.HTML(200, "data/index.html", gin.H{"messages": rows})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("nick")
	session.Save()
	c.Redirect(302, "/")
}

func ws(c *gin.Context) {
	session := sessions.Default(c)

	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	conns = append(conns, conn)

	go func() {
		for {
			t, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}

			nick := session.Get("nick")

			Insert(nick.(string), string(msg), time.Now().String())

			for _, c := range conns {
				c.WriteMessage(t, []byte(fmt.Sprintf("%s:%s", nick, msg)))
			}
		}
	}()
}

type MyHTMLRender struct{}

func (r *MyHTMLRender) Instance(name string, data interface{}) render.Render {
	tmpl, _ := Asset(name)
	t := template.New(name)
	t.Parse(string(tmpl))
	return render.HTML{
		Template: t,
		Data:     data,
	}
}

func main() {
	Init()
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()
	render := &MyHTMLRender{}

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.HTMLRender = render

	r.GET("/", index)
	r.GET("/logout", logout)
	r.GET("/ws", ws)

	r.Run(":8080")
}
