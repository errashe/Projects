package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"

	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	conns []*websocket.Conn
)

func error_handler(err error) {
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
}

func broadcast(msg []byte) {
	for _, c := range conns {
		broadcastOne(c, msg)
	}
}

func broadcastOne(conn *websocket.Conn, msg []byte) {
	conn.WriteMessage(1, msg)
}

func loadTemplates() *template.Template {
	var templates = template.New("")

	funcMap := template.FuncMap{
		"isSessionIsset": func(c *http.Request) bool {
			session, _ := store.Get(c, "main")
			if session.Values["nick"] == nil {
				return false
			}
			return true
		},
	}

	templates.Funcs(funcMap)

	for _, path := range AssetNames() {
		bytes, err := Asset(path)
		error_handler(err)
		templates.New(path).Parse(string(bytes))
	}

	return templates
}

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

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(loadTemplates())

	initDB()
	defer db.Close()

	r.GET("/", main_handler)
	r.POST("/login", login_handler)
	r.GET("/logout", logout_handler)
	r.GET("/ws", ws_handler)

	go func() {
		for {
			var str string
			fmt.Scanln(&str)
			s := strings.Split(str, ",")

			switch s[0] {
			case "au":
				addUser(s[1], s[2])
			case "am":
				addMessage(s[1], s[2])
			case "gm":
				id, err := strconv.Atoi(s[1])
				error_handler(err)
				fmt.Printf("%v\n", getUser(id))
			case "cu":
				fmt.Println(checkUser(s[1], s[2]))
			}
		}
	}()

	r.Run()
}
